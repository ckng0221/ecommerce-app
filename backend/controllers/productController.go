package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
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

func ConsumeProductStock() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{
				"message": "Invalid request body",
			})
			return
		}

		var stockUpdate = struct {
			Quantity int `json:"stock_quantity"`
		}{}

		err = json.Unmarshal(body, &stockUpdate)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{
				"message": "failed to parse request body",
			})
			return
		}
		if stockUpdate.Quantity <= 0 {
			render.Status(r, 422)
			render.JSON(w, r, map[string]string{
				"message": "Quantity cannot less than or equal to 0",
			})
			return
		}

		var product models.Product

		initializers.Db.First(&product, id)

		if product.StockQuantity < stockUpdate.Quantity {
			render.Status(r, 422)
			render.JSON(w, r, map[string]string{
				"message": "insufficient stock",
			})
			return
		}

		expression := "stock_quantity - ?"

		result := initializers.Db.Model(&product).UpdateColumn("stock_quantity", gorm.Expr(expression, stockUpdate.Quantity))
		initializers.Db.First(&product, product.ID)

		if result.Error != nil {
			fmt.Println(result.Error)
			render.Status(r, 400)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, map[string]int{
			"stock_quantity": product.StockQuantity,
		})
	}
}

func AddProductStock() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{
				"message": "Invalid request body",
			})
			return
		}

		var stockUpdate = struct {
			Quantity int `json:"stock_quantity"`
		}{}

		err = json.Unmarshal(body, &stockUpdate)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{
				"message": "failed to parse request body",
			})
			return
		}
		if stockUpdate.Quantity <= 0 {
			render.Status(r, 422)
			render.JSON(w, r, map[string]string{
				"message": "Quantity cannot less than or equal to 0",
			})
			return
		}

		var product models.Product

		initializers.Db.First(&product, id)

		expression := "stock_quantity + ?"

		result := initializers.Db.Model(&product).UpdateColumn("stock_quantity", gorm.Expr(expression, stockUpdate.Quantity))
		initializers.Db.First(&product, product.ID)

		if result.Error != nil {
			fmt.Println(result.Error)
			render.Status(r, 400)
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, map[string]int{
			"stock_quantity": product.StockQuantity,
		})
	}
}
