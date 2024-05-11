package controllers

import (
	"ecommerce-app/models"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	order := []struct {
		ID     uint
		UserID string `json:"user_id"`
	}{}
	GetAllNew[models.Order](w, r, order)

}

func CreateOrders() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Order]
}

func GetOrderById() func(w http.ResponseWriter, r *http.Request) {
	return GetById[models.Order]
}

func UpdateOrderById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Order, models.OrderUpdate]
}

func DeleteOrderById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Order]
}
