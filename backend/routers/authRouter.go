package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", controllers.Login)
	r.Get("/google-login", controllers.GoogleLogin)

	r.With(middlewares.RequireAuth).Get("/validate", controllers.Validate)
	return r
}
