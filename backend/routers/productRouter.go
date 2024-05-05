package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi"
)

func ProductRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetProducts())

	r.Post("/", controllers.CreateProducts())

	r.Get("/{id}", controllers.GetProductById())

	r.Patch("/{id}", controllers.UpdateProductById())

	r.Delete("/{id}", controllers.DeleteProductById())

	r.Post("/{id}/stock", controllers.UpdateProductStock)

	return r
}
