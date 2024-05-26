package controllers

import (
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"net/http"

	"clevergo.tech/jsend"
	"gorm.io/gorm"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	scope := utils.EmptyScope

	userId := r.URL.Query().Get("user_id")
	if userId != "" {
		err := requireOwner(r, userId)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
		scope = func(db *gorm.DB) *gorm.DB {
			return db.Preload("Address").Preload("OrderItems.Product").Where("user_id = ?", userId).Order("id desc")
		}
	} else {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	}
	GetAll[models.Order](w, r, scope)
}

func CreateOrders() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Order]
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	var scope = func(db *gorm.DB) *gorm.DB {
		return db.Preload("OrderItems.Product").Preload("User").Preload("Address")
	}
	GetById[models.Order](w, r, scope)
}

func UpdateOrderById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Order, models.OrderUpdate]
}

func DeleteOrderById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Order]
}
