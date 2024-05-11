package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi"
)

func OrderRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetOrders)

	r.Post("/", controllers.CreateOrders())

	r.Get("/{id}", controllers.GetOrderById())

	r.Patch("/{id}", controllers.UpdateOrderById())

	r.Delete("/{id}", controllers.DeleteOrderById())

	return r
}
