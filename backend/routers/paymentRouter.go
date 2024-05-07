package routers

import (
	"ecommerce-app/controllers"

	"github.com/go-chi/chi"
)

func PaymentRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/checkout/session", controllers.CreateCheckoutSession)
	r.Post("/webhook", controllers.StripePaymentHook)
	r.Post("/event/trigger", controllers.TriggerFakePaymentEvent)

	return r
}
