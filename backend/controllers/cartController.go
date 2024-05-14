package controllers

import (
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"net/http"

	"gorm.io/gorm"
)

func GetCarts(w http.ResponseWriter, r *http.Request) {
	scope := utils.EmptyScope

	userId := r.URL.Query().Get("user_id")
	if userId != "" {
		scope = func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userId)
		}
	}

	GetAll[models.Cart](w, r, scope)
}

func CreateCarts() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Cart]
}

func GetCartById(w http.ResponseWriter, r *http.Request) {
	GetById[models.Cart](w, r, utils.EmptyScope)
}

func UpdateCartById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Cart, models.CartUpdate]
}

func DeleteCartById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Cart]
}
