package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func UserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.With(middlewares.RequireAuth).Get("/", controllers.GetUsers)

	r.With(middlewares.RequireAuth).Post("/", controllers.CreateUser)

	r.Post("/signup", controllers.Signup)

	r.With(middlewares.RequireAuth).Get("/{id}", controllers.GetUserById)
	r.With(middlewares.RequireAuth).Get("/sub/{sub}", controllers.GetUserBySub)

	r.With(middlewares.RequireAuth).Get("/{id}/addresses", controllers.GetAddressesByUserId)

	r.With(middlewares.RequireAuth).Post("/{id}/addresses", controllers.CreateAddressByUserId)

	r.With(middlewares.RequireAuth).Patch("/{id}", controllers.UpdateUserById)

	r.With(middlewares.RequireAuth).Delete("/{id}", controllers.DeleteUserById)

	// Address
	r.With(middlewares.RequireAuth).Get("/addresses/{id}", controllers.GetAddressById)
	r.With(middlewares.RequireAuth).Post("/addresses", controllers.CreateAddress)
	r.With(middlewares.RequireAuth).Patch("/addresses/{id}", controllers.UpdateAddressById)
	r.With(middlewares.RequireAuth).Delete("/addresses/{id}", controllers.DeleteAddressById)

	return r
}
