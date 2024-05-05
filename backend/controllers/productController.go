package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/go-chi/chi"
)

func GetProducts() func(w http.ResponseWriter, r *http.Request) {
	return GetAll[models.Product]
}

func CreateProducts() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Product]
}

func GetProductById() func(w http.ResponseWriter, r *http.Request) {
	return GetById[models.Product]
}

func UpdateProductById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Product, models.ProductUpdate]
}

func DeleteProductById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Product]
}

func UpdateProductStockCount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid request body"))
		return
	}

	var stockUpdate = struct {
		Action   string `json:"action"`
		Quantity int    `json:"stock_quantity"`
	}{}

	err = json.Unmarshal(body, &stockUpdate)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("failed to parse request body"))
		return
	}
	var availableActions = []string{"add", "consume"}
	if !slices.Contains(availableActions, stockUpdate.Action) {
		w.WriteHeader(422)
		w.Write([]byte("Invalid action"))
		return
	}

	var product models.Product

	initializers.Db.First(&product, id)

	if stockUpdate.Action == "consume" && product.StockQuantity < stockUpdate.Quantity {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "insufficient stock",
		})
		return
	}

	finalQuantity := product.StockQuantity
	switch stockUpdate.Action {
	case "add":
		finalQuantity = product.StockQuantity + stockUpdate.Quantity
	case "consume":
		finalQuantity = product.StockQuantity - stockUpdate.Quantity
	}

	// Optimistic lock
	result := initializers.Db.Model(&product).Where("id = ? AND stock_quantity = ?",
		id, product.StockQuantity).Update("stock_quantity", finalQuantity)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(product)
}
