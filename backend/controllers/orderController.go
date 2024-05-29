package controllers

import (
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"errors"
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
	var order models.Order

	GetById(w, r, scope, &order, true)
}

func UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	err := requireAdmin(r)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	var order models.Order
	var orderUpdate models.OrderUpdate

	UpdateById(w, r, utils.EmptyScope, &order, &orderUpdate, false)
}

func DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	err := requireAdmin(r)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	DeleteById[models.Order](w, r)
}
