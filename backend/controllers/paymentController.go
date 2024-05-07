package controllers

import (
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/render"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/webhook"
)

type CheckoutItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func productIdExists(checkoutItems []models.Product, productID uint) bool {
	for _, product := range checkoutItems {
		if product.ID == productID {
			return true
		}
	}
	return false
}

func (c *CheckoutItem) validate() url.Values {
	errs := url.Values{}
	if c.Quantity <= 0 {
		errs.Add("quantity", "Quantity has to be greater than 0")
	}
	return errs
}

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	var checkoutItems []CheckoutItem
	var products []models.Product
	var idList []uint

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request body",
		})
		return
	}

	err = json.Unmarshal(body, &checkoutItems)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse request body",
		})
		return
	}
	// Process request body
	for _, item := range checkoutItems {
		if validErrs := item.validate(); len(validErrs) > 0 {
			err := map[string]interface{}{"message": validErrs}
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err)
			return
		}
		idList = append(idList, item.ProductID)
	}

	err = initializers.Db.Find(&products, idList).Error
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	// Check if all proudcts can be found
	for _, productId := range idList {
		if !productIdExists(products, productId) {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf("Product ID: %v not found", productId),
			})
			return
		}
	}

	var frontend_base_url = os.Getenv("FRONTEND_BASE_URL")
	stripe.Key = os.Getenv("STRIPE_KEY")
	successUrl := fmt.Sprintf("%s/successful", frontend_base_url)
	cancelUrl := fmt.Sprintf("%s/cancel", frontend_base_url)
	lineItems := []*stripe.CheckoutSessionLineItemParams{}

	// Convert to stripe slice
	for idx, product := range products {
		var itemQuantity int = 1
		if checkoutItems[idx].ProductID == product.ID {
			itemQuantity = checkoutItems[idx].Quantity
		}

		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String(product.Currency),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:        stripe.String(product.Name),
					Description: stripe.String(product.Description),
				},
				UnitAmount: stripe.Int64(int64(product.UnitPrice * 100)),
			},
			Quantity: stripe.Int64(int64(itemQuantity)),
		})
	}

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

func StripePaymentHook(w http.ResponseWriter, r *http.Request) {
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
