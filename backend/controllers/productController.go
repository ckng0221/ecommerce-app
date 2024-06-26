package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"clevergo.tech/jsend"
	"gorm.io/gorm"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	GetAll[models.Product](w, r, utils.EmptyScope)
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	CreateOne(w, r, &product, true, false)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	GetById(w, r, utils.EmptyScope, &product, false, false)
}

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	var orderUpdate models.ProductUpdate

	UpdateById(w, r, utils.EmptyScope, &product, &orderUpdate, true, false)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	DeleteById(w, r, utils.EmptyScope, &product, false, true)
}

func ConsumeProductStock() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var stockUpdate = struct {
			Quantity int `json:"stock_quantity"`
		}{}

		err = json.Unmarshal(body, &stockUpdate)
		if err != nil {
			jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
			return
		}
		if stockUpdate.Quantity <= 0 {
			jsend.Fail(w, "Quantity cannot less than or equal to 0", http.StatusUnprocessableEntity)
			return
		}

		var product models.Product

		initializers.Db.First(&product, id)

		if product.StockQuantity < stockUpdate.Quantity {
			jsend.Fail(w, "insufficient stock", http.StatusBadRequest)
			return
		}

		expression := "stock_quantity - ?"
		result := initializers.Db.Model(&product).UpdateColumn("stock_quantity", gorm.Expr(expression, stockUpdate.Quantity))

		if result.Error != nil {
			log.Println(result.Error)
			jsend.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		initializers.Db.First(&product, product.ID)

		jsend.Success(w, map[string]int{"stock_quantity": product.StockQuantity})
	}
}

func AddProductStock() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var stockUpdate = struct {
			Quantity int `json:"stock_quantity"`
		}{}

		err = json.Unmarshal(body, &stockUpdate)
		if err != nil {
			jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
			return
		}
		if stockUpdate.Quantity <= 0 {
			jsend.Fail(w, "Quantity cannot less than or equal to 0", http.StatusUnprocessableEntity)
			return
		}

		var product models.Product

		initializers.Db.First(&product, id)

		expression := "stock_quantity + ?"

		result := initializers.Db.Model(&product).UpdateColumn("stock_quantity", gorm.Expr(expression, stockUpdate.Quantity))
		initializers.Db.First(&product, product.ID)

		if result.Error != nil {
			log.Println(result.Error)
			jsend.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		jsend.Success(w, map[string]int{"stock_quantity": product.StockQuantity})
	}
}
