package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"clevergo.tech/jsend"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	GetAll[models.User](w, r, utils.EmptyScope)
}

func CreateUser() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.User]
}

func CreateUserAddress() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Address]
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var scope = func(db *gorm.DB) *gorm.DB {
		return db.Preload("DefaultAddress")
	}

	GetById[models.User](w, r, scope)
}

func GetAddressesByUserId(w http.ResponseWriter, r *http.Request) {
	GetChildrenById[models.Address](w, r, "address_id", "User")
}

func CreateAddressByUserId(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	id := chi.URLParam(r, "id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}
	uID, _ := strconv.ParseUint(id, 10, 32)
	address.UserID = uint(uID)
	fmt.Println(address)

	result := initializers.Db.Model(&models.Address{}).Create(&address)
	if result.Error != nil {
		fmt.Println(result.Error)

		jsend.Error(w, "failed to create item", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, address, http.StatusCreated)
}

func UpdateUserById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.User, models.UserUpdate]
}

func DeleteUserById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.User]
}
