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
	"strconv"

	"clevergo.tech/jsend"
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
	err := requireAdmin(r)
	if err != nil {
		jsend.Fail(w, "Forbidden", http.StatusForbidden)
		return
	}
	GetAll[models.User](w, r, utils.EmptyScope)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := requireAdmin(r)
	if err != nil {
		jsend.Fail(w, "Forbidden", http.StatusForbidden)
		return
	}

	CreateOne[models.User](w, r)
}

func CreateUserAddress() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Address]
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User

	scope := utils.EmptyScope
	scope = func(db *gorm.DB) *gorm.DB {
		return db.Joins("DefaultAddress")
	}

	GetById(w, r, scope, &user, true)
}

func GetUserBySub(w http.ResponseWriter, r *http.Request) {
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Joins("DefaultAddress")
	}

	sub := r.PathValue("sub")

	var user models.User

	err := initializers.Db.Scopes(scope).Where("sub = ?", sub).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusBadRequest)
			return
		}
		log.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = requireOwner(r, fmt.Sprint(user.ID))

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

	jsend.Success(w, &user)
}

func GetAddressesByUserId(w http.ResponseWriter, r *http.Request) {
	GetChildrenById[models.Address](w, r, "user_id", "User")
}

func CreateAddressByUserId(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	id := r.PathValue("id")

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

	// Require Owner
	err = requireOwner(r, fmt.Sprint(uID))
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

	result := initializers.Db.Model(&models.Address{}).Create(&address)
	if result.Error != nil {
		log.Println(result.Error)

		jsend.Error(w, "failed to create item", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, address, http.StatusCreated)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userUpdate models.UserUpdate

	UpdateById(w, r, utils.EmptyScope, &user, &userUpdate, true)
}

func DeleteUserById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.User]
}

func GetAddressById(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	scope := utils.EmptyScope

	GetById(w, r, scope, &address, true)
}

func UpdateAddressById(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	var addressUpdate models.AddressUpdate

	UpdateById(w, r, utils.EmptyScope, &address, &addressUpdate, true)
}

func DeleteAddressById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Address]
}

func CreateAddress() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Address]
}
