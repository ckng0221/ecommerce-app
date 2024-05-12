package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"clevergo.tech/jsend"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
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

func CreateOne[T Model](w http.ResponseWriter, r *http.Request) {
	var modelObj T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &modelObj)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	result := initializers.Db.Model(&modelObj).Create(&modelObj)
	if result.Error != nil {
		fmt.Println(result.Error)

		jsend.Error(w, "failed to create item", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, &modelObj, http.StatusCreated)
}

func GetById[T Model](w http.ResponseWriter, r *http.Request, scope func(db *gorm.DB) *gorm.DB) {
	id := chi.URLParam(r, "id")

	var modelObj T

	err := initializers.Db.Scopes(scope).First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, &modelObj)
}

func GetChildrenById[T Model](w http.ResponseWriter, r *http.Request, chilrenIdName string, preloadName string) {
	var modelObjs []T
	id := chi.URLParam(r, "id")

	paginationScope, error := utils.Paginate(r)
	if error != nil {
		jsend.Fail(w, error.Error(), http.StatusBadRequest)
		return
	}

	expression := fmt.Sprintf("%s = ?", chilrenIdName)
	initializers.Db.Preload(preloadName).Scopes(paginationScope).Find(&modelObjs).Where(expression, id)
	jsend.Success(w, &modelObjs)
}

func UpdateById[T Model, TUpdate Model](w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var modelObj T
	var modelUpdateObj TUpdate

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &modelObj)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	err = initializers.Db.First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusNotFound)
			return
		}
		fmt.Println(err)
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	initializers.Db.Model(&modelObj).Updates(&modelUpdateObj)

	jsend.Success(w, &modelObj)

}

func DeleteById[T Model](w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var modelObj T

	err := initializers.Db.First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			jsend.Fail(w, "Record not found", http.StatusNotFound)
			return

		}
		fmt.Println(err)

		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result := initializers.Db.Delete(&modelObj, id)
	if result.Error != nil {
		fmt.Println(err)

		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, nil, http.StatusNoContent)
}
