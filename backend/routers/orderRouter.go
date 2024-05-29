package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func OrderRouter() *chi.Mux {
	r := chi.NewRouter()

	r.With(middlewares.RequireAuth).Get("/", controllers.GetOrders)

	r.With(middlewares.RequireAuth).Post("/", controllers.CreateOrders())

	// r.Get("/{id}", controllers.GetOrderById)
	r.With(middlewares.RequireAuth).Get("/{id}", controllers.GetOrderById)

	r.With(middlewares.RequireAuth).Patch("/{id}", controllers.UpdateOrderById)

	r.With(middlewares.RequireAuth).Delete("/{id}", controllers.DeleteOrderById)

	return r
}
