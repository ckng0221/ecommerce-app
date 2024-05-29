package routers

import (
	"ecommerce-app/controllers"
	"ecommerce-app/middlewares"

	"github.com/go-chi/chi/v5"
)

func CartRouter() *chi.Mux {
	r := chi.NewRouter()

	r.With(middlewares.RequireAuth).Get("/", controllers.GetCarts)

	r.With(middlewares.RequireAuth).Post("/", controllers.CreateCarts())

	r.With(middlewares.RequireAuth).Post("/add", controllers.CreateOrAddCart)

	r.With(middlewares.RequireAuth).Get("/{id}", controllers.GetCartById)

	r.With(middlewares.RequireAuth).Patch("/{id}", controllers.UpdateCartById)

	r.With(middlewares.RequireAuth).Delete("/{id}", controllers.DeleteCartById())

	return r
}
