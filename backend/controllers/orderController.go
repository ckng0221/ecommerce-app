package controllers

import (
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order

	GetAllNew[models.Order](w, r, orders, utils.NullScope)

}

func CreateOrders() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Order]
}

func GetOrderById() func(w http.ResponseWriter, r *http.Request) {
	// var scope = func(db *gorm.DB) *gorm.DB {
	// 	return db.Preload("OrderItems").Preload("User").Preload("Address")
	// }
	return GetById[models.Order]
}

func UpdateOrderById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Order, models.OrderUpdate]
}

func DeleteOrderById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Order]
}
