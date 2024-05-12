package controllers

import (
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"net/http"

	"gorm.io/gorm"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {

	GetAll[models.Order](w, r, utils.EmptyScope)

}

// TODO: need to create a page in frontend for creating order first, then only proceed to payment
func CreateOrders() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Order]
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	var scope = func(db *gorm.DB) *gorm.DB {
		return db.Preload("OrderItems").Preload("User").Preload("Address")
	}
	GetById[models.Order](w, r, scope)
}

func UpdateOrderById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Order, models.OrderUpdate]
}

func DeleteOrderById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Order]
}
