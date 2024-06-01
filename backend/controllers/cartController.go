package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"clevergo.tech/jsend"
	"gorm.io/gorm"
)

func GetCarts(w http.ResponseWriter, r *http.Request) {
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product")
	}

	userId := r.URL.Query().Get("user_id")
	if userId != "" {
		err := requireOwner(r, userId)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
		scope = func(db *gorm.DB) *gorm.DB {
			return db.Preload("Product").Where("user_id = ?", userId)
		}
	} else {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	GetAll[models.Cart](w, r, scope)
}

func CreateCarts(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	CreateOne(w, r, &cart, false, true)
}

func CreateOrAddCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &cart)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// Require owner
	err = requireOwner(r, fmt.Sprint(cart.UserID))
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			log.Println(err.Error())
			jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Check if cart with productId exists
	var existingCart models.Cart
	initializers.Db.Model(&models.Cart{}).Where("user_id = ? AND product_id = ?", cart.UserID, cart.ProductID).Find(&existingCart)
	if existingCart.ID != 0 {
		existingCart.Quantity += cart.Quantity
		result := initializers.Db.Updates(&existingCart)
		if result.Error != nil {
			log.Println(result.Error)

			jsend.Error(w, "failed to update item", http.StatusInternalServerError)
			return
		}

		jsend.Success(w, &existingCart)
	} else {
		result := initializers.Db.Model(&models.Cart{}).Create(&cart)
		if result.Error != nil {
			log.Println(result.Error)

			jsend.Error(w, "failed to create item", http.StatusInternalServerError)
			return
		}

		jsend.Success(w, &cart, http.StatusCreated)
	}
}

func GetCartById(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart

	GetById(w, r, utils.EmptyScope, &cart, false, true)
}

func UpdateCartById(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	var cartUpdate models.CartUpdate

	UpdateById(w, r, utils.EmptyScope, &cart, &cartUpdate, false, true)
}

func DeleteCartById(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	DeleteById(w, r, utils.EmptyScope, &cart, false, true)
}
