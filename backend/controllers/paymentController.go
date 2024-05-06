package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/webhook"
)

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	var frontend_base_url = os.Getenv("FRONTEND_BASE_URL")
	stripe.Key = os.Getenv("STRIPE_KEY")
	successUrl := fmt.Sprintf("%s/successful", frontend_base_url)
	cancelUrl := fmt.Sprintf("%s/cancel", frontend_base_url)
	lineItems := []*stripe.CheckoutSessionLineItemParams{}

	lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("myr"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name:        stripe.String("Apple Watch"),
				Description: stripe.String("This is an Apple Watch"),
			},
			UnitAmount: stripe.Int64(100000),
		},
		Quantity: stripe.Int64(int64(1)),
	})

	params := &stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String(successUrl),
		CancelURL:          stripe.String(cancelUrl),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	result, err := session.New(params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	render.JSON(w, r, map[string]string{"url": result.URL})
}

func StipePaymentHook(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	endpointSecret := os.Getenv("STRIPE_CLI_WEBHOOK_SECRET")
	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	// fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	fmt.Printf("Unhandled event type: %s\n", event.Type)

	w.WriteHeader(http.StatusOK)
}
