package controllers

import (
	"bytes"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"clevergo.tech/jsend"
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
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	err = json.Unmarshal(body, &checkoutItems)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Process request body
	for _, item := range checkoutItems {
		if validErrs := item.validate(); len(validErrs) > 0 {
			err := map[string]interface{}{"message": validErrs}
			jsend.Fail(w, err, http.StatusBadRequest)
			return
		}
		idList = append(idList, item.ProductID)
	}

	err = initializers.Db.Find(&products, idList).Error
	if err != nil {
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Check if all proudcts can be found
	for _, productId := range idList {
		if !productIdExists(products, productId) {
			jsend.Fail(w, fmt.Sprintf("Product ID: %v not found", productId), http.StatusBadRequest)

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
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, map[string]string{"url": result.URL})
}

// Trigger fake payment event with stripe CLI (on local only)
//
// In dev environment with public URL, could put actual endpoitn.
func TriggerFakePaymentEvent(w http.ResponseWriter, r *http.Request) {
	type Event struct {
		Event string `json:"event"`
	}
	body, _ := io.ReadAll(r.Body)

	var event Event
	// default event
	var userEvent string = "payment_intent.succeeded"
	if len(body) > 0 {
		err := json.Unmarshal(body, &event)
		if err != nil {
			jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		}
		userEvent = event.Event

	}

	cmd := exec.Command("stripe", "trigger", userEvent)
	var out bytes.Buffer
	// define the process standard output
	cmd.Stdout = &out
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(out.String())

	jsend.Success(w, map[string]string{"message": "Triggered payment event"}, http.StatusAccepted)
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
	// fmt.Println("signature", r.Header.Get("Stripe-Signature"))

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		jsend.Fail(w, nil, http.StatusBadRequest)
		return
	}

	fmt.Printf("Unhandled event type: %s\n", event.Type)

	// TODO: update order for payment done

	jsend.Success(w, nil, http.StatusOK)
}
