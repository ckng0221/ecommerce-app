package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"clevergo.tech/jsend"
	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Manual signup, without social login
func Signup(w http.ResponseWriter, r *http.Request) {
	// Get the email/pass off req bodyu
	var requestBody models.UserSignUp

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate email
	var checkUser models.User
	initializers.Db.Model(&models.User{}).Where("email = ?", requestBody.Email).First(&checkUser)

	if checkUser.ID != 0 {
		jsend.Fail(w, "email is already in use", http.StatusBadRequest)
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)

	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// Create the user
	password := string(hash)
	user := models.User{Name: requestBody.Name, Email: requestBody.Email, Password: &password, Role: "member"}
	result := initializers.Db.Create(&user)

	if result.Error != nil {
		log.Println("failed to create user")
		log.Println(result.Error)

		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, &user, http.StatusCreated)
}

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
	log.Println(address)

	result := initializers.Db.Model(&models.Address{}).Create(&address)
	if result.Error != nil {
		log.Println(result.Error)

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
