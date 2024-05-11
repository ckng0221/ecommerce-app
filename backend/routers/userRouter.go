package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi"
)

func UserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetUsers())

	r.Post("/", controllers.CreateUser())

	r.Get("/{id}", controllers.GetUserById())

	r.Get("/{id}/addresses", controllers.GetAddressesByUserId)

	r.Post("/{id}/addresses", controllers.CreateAddressByUserId)

	r.Patch("/{id}", controllers.UpdateUserById())

	r.Delete("/{id}", controllers.DeleteUserById())

	return r
}
