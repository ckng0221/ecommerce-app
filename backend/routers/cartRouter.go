package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi"
)

func CartRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetCarts)

	r.Post("/", controllers.CreateCarts())

	r.Get("/{id}", controllers.GetCartById)

	r.Patch("/{id}", controllers.UpdateCartById())

	r.Delete("/{id}", controllers.DeleteCartById())

	return r
}
