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
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Model interface{}

func GetAll[T Model](w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB) {
	var modelObjs []T

	paginationScope, error := utils.Paginate(r)
	if error != nil {
		jsend.Fail(w, error.Error(), http.StatusBadRequest)
		return
	}

	initializers.Db.Model(&modelObjs).Scopes(paginationScope, scope).Find(&modelObjs)
	jsend.Success(w, &modelObjs)
}

func CreateOne(w http.ResponseWriter, r *http.Request, modelObj interface{}, needAdmin, needOwner bool) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, modelObj)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	if needAdmin {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else if needOwner {
		var userId uint = 0
		switch v := (modelObj).(type) {
		case *models.Cart:
			userId = v.UserID
		case *models.Address:
			userId = v.UserID
		default:
			log.Println("Cannot match type")
		}
		err := requireOwner(r, fmt.Sprint(userId))

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
	}

	result := initializers.Db.Model(modelObj).Create(modelObj)
	if result.Error != nil {
		log.Println(result.Error)

		jsend.Error(w, "failed to create item", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, modelObj, http.StatusCreated)
}

func GetById(w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB, modelObj interface{}, needAdmin, needOwner bool) {
	id := r.PathValue("id")

	if id == "" {
		log.Println("id is empty")
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := initializers.Db.Scopes(scope).First(modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusBadRequest)
			return
		}
		log.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if needAdmin {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else if needOwner {
		var userId uint = 0
		switch v := (modelObj).(type) {
		case *models.Order:
			userId = v.UserID
		case *models.Cart:
			userId = v.UserID
		case *models.Address:
			userId = v.UserID
		case *models.User:
			userId = v.ID
		default:
			log.Println("Cannot match type")
		}
		err := requireOwner(r, fmt.Sprint(userId))

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
	}

	jsend.Success(w, modelObj)
}

func GetChildrenById(w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB, modelObjs interface{}, chilrenIdName string, needAdmin, needOwner bool) {
	id := r.PathValue("id")

	if needAdmin {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else if needOwner {
		var userId uint = 0
		switch (modelObjs).(type) {
		case *[]models.Address:
			userId64, _ := strconv.ParseUint(id, 10, 32)
			userId = uint(userId64)
		default:
			log.Println("Cannot match type")
		}
		err := requireOwner(r, fmt.Sprint(userId))

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
	}

	paginationScope, err := utils.Paginate(r)
	if err != nil {
		jsend.Fail(w, err.Error(), http.StatusBadRequest)
		return
	}

	expression := fmt.Sprintf("%s = ?", chilrenIdName)
	initializers.Db.Scopes(paginationScope, scope).Where(expression, id).Find(modelObjs)
	jsend.Success(w, modelObjs)
}

func UpdateById(w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB, modelObj, modelUpdateObj interface{}, needAdmin, needOwner bool) {
	id := r.PathValue("id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, modelUpdateObj)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	err = initializers.Db.First(modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusNotFound)
			return
		}
		log.Println(err)
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if needAdmin {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else if needOwner {
		// Note: those require admin, don't require owner
		var userId uint = 0
		switch v := (modelObj).(type) {

		case *models.Cart:
			userId = v.UserID
		case *models.Address:
			userId = v.UserID
		case *models.User:
			userId = v.ID
		default:
			log.Println("Cannot match type")
		}
		// fmt.Println("lalal", userId)
		err := requireOwner(r, fmt.Sprint(userId))

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
	}

	initializers.Db.Clauses(clause.Returning{}).Model(modelObj).Updates(modelUpdateObj)
	jsend.Success(w, modelObj)
}

func DeleteById(w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB, modelObj interface{}, needAdmin, needOwner bool) {
	id := r.PathValue("id")

	err := initializers.Db.First(modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusNotFound)
			return

		}
		log.Println(err)

		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if needAdmin {
		err := requireAdmin(r)
		if err != nil {
			jsend.Fail(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else if needOwner {
		// Note: those require admin, don't require owner
		var userId uint = 0
		switch v := (modelObj).(type) {

		case *models.Cart:
			userId = v.UserID
		case *models.Address:
			userId = v.UserID
		case *models.User:
			userId = v.ID
		default:
			log.Println("Cannot match type")
		}
		// fmt.Println("lalal", userId)
		err := requireOwner(r, fmt.Sprint(userId))

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
	}

	result := initializers.Db.Delete(modelObj, id)
	if result.Error != nil {
		log.Println(err)

		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, nil, http.StatusNoContent)
}
