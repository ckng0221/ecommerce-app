package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi/v5"
)

func PaymentRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.GetPayments)
	r.Post("/checkout/session", controllers.CreatePaymentSession)
	r.Post("/webhook", controllers.StripePaymentHook)
	r.Post("/event/trigger", controllers.TriggerFakePaymentEvent)

	return r
}
