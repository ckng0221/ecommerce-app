package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func UserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetUsers)

	r.Post("/", controllers.CreateUser())

	r.Post("/signup", controllers.Signup)

	r.With(middlewares.RequireAuth).Get("/{id}", controllers.GetUserById)
	r.Get("/sub/{sub}", controllers.GetUserBySub)

	r.Get("/{id}/addresses", controllers.GetAddressesByUserId)

	r.Post("/{id}/addresses", controllers.CreateAddressByUserId)

	r.Patch("/{id}", controllers.UpdateUserById())

	r.Delete("/{id}", controllers.DeleteUserById())

	// Address
	r.Get("/addresses/{id}", controllers.GetAddressById)
	r.Post("/addresses", controllers.CreateAddress())
	r.Patch("/addresses/{id}", controllers.UpdateAddressById())
	r.Delete("/addresses/{id}", controllers.DeleteAddressById())

	return r
}
