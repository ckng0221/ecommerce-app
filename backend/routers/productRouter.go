package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func ProductRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetProducts)

	r.With(middlewares.RequireAuth).Post("/", controllers.CreateProducts)

	r.Get("/{id}", controllers.GetProductById)

	r.With(middlewares.RequireAuth).Patch("/{id}", controllers.UpdateProductById)

	r.With(middlewares.RequireAuth).Delete("/{id}", controllers.DeleteProductById)

	r.Post("/{id}/stocks/consume", controllers.ConsumeProductStock())

	r.Post("/{id}/stocks/add", controllers.AddProductStock())

	return r
}
