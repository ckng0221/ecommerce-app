package controllers

import (
	"ecommerce-app/models"
	"net/http"
)

func GetProducts() func(w http.ResponseWriter, r *http.Request) {
	return GetAll[models.Product]
}

func CreateProducts() func(w http.ResponseWriter, r *http.Request) {
	return CreateOne[models.Product]
}

func GetProductById() func(w http.ResponseWriter, r *http.Request) {
	return GetById[models.Product]
}

func UpdateProductById() func(w http.ResponseWriter, r *http.Request) {
	return UpdateById[models.Product, models.ProductUpdate]
}

func DeleteProductById() func(w http.ResponseWriter, r *http.Request) {
	return DeleteById[models.Product]
}
