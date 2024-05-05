package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type Model interface{}

func GetAll[T Model](w http.ResponseWriter, r *http.Request) {
	var modelObjs []T

	paginationScope, error := utils.Paginate(r)
	if error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": error.Error(),
		})
		return
	}

	initializers.Db.Scopes(paginationScope).Find(&modelObjs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&modelObjs)
}

func CreateOne[T Model](w http.ResponseWriter, r *http.Request) {
	var modelObj T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	err = json.Unmarshal(body, &modelObj)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse request body",
		})
		return
	}

	result := initializers.Db.Model(&modelObj).Create(&modelObj)
	if result.Error != nil {
		fmt.Println(result.Error)

		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to create item",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&modelObj)
}

func GetById[T Model](w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var modelObj T

	err := initializers.Db.First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Record not found",
			})
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&modelObj)
}

func UpdateById[T Model, TUpdate Model](w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var modelObj T
	var modelUpdateObj TUpdate

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	err = json.Unmarshal(body, &modelObj)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Failed to parse request body",
		})
		return
	}

	err = initializers.Db.First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Record not found",
			})
			return
		}
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}
	initializers.Db.Model(&modelObj).Updates(&modelUpdateObj)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&modelObj)
}

func DeleteById[T Model](w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var modelObj T

	err := initializers.Db.First(&modelObj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Record not found",
			})
			return
		}
		fmt.Println(err)

		w.WriteHeader(500)
		return
	}

	result := initializers.Db.Delete(&modelObj, id)
	if result.Error != nil {
		fmt.Println(err)

		w.WriteHeader(500)
		return
	}

	w.WriteHeader(202)
}
