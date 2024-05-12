package controllers

import (
	"bytes"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

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

// Create Payment session in Stripe
func CreatePaymentSession(w http.ResponseWriter, r *http.Request) {
	type CheckoutRequestBody struct {
		OrderID       uint           `json:"order_id"`
		CheckoutItems []CheckoutItem `json:"checkout_items"`
	}

	var checkoutRequestBody CheckoutRequestBody
	var products []models.Product
	var idList []uint

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	err = json.Unmarshal(body, &checkoutRequestBody)
	if err != nil {
		jsend.Fail(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	checkoutItems := checkoutRequestBody.CheckoutItems

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
		log.Println(err)
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

	params.AddMetadata("order_id", fmt.Sprint(checkoutRequestBody.OrderID))

	// log.Println("param", params)
	result, err := session.New(params)
	if err != nil {
		log.Println(err)
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create Payment record
	var payment = models.Payment{
		OrderID:         checkoutRequestBody.OrderID,
		StripeSessionID: result.ID,
	}
	err = initializers.Db.Model(&models.Payment{}).Create(&payment).Error
	if err != nil {
		log.Println("Failed to create payment in db")
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
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	log.Println(out.String())

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
	// log.Println("signature", r.Header.Get("Stripe-Signature"))

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		jsend.Fail(w, nil, http.StatusBadRequest)
		return
	}

	err = processPaymentEvent(event)

	if err != nil {
		log.Println(err.Error())
		jsend.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsend.Success(w, nil, http.StatusOK)
}

func processPaymentEvent(event stripe.Event) error {
	switch event.Type {
	case "checkout.session.completed":
		var checkoutSession stripe.CheckoutSession
		// log.Println("raw", string(event.Data.Raw))
		err := json.Unmarshal(event.Data.Raw, &checkoutSession)
		if err != nil {
			log.Println("Failed to parse event data raw.")
			return err
		}

		orderID := checkoutSession.Metadata["order_id"]
		stripeCheckoutSessionID := checkoutSession.ID

		log.Printf("Received payment for order ID: %v\n", orderID)

		// Further processing based on metadata...
		var order models.Order
		err = initializers.Db.First(&order, orderID).Error
		if err != nil {
			log.Println(err)
			return err
		}

		currentTime := time.Now()
		var orderUpdate = models.OrderUpdate{
			PaymentAt:   &currentTime,
			OrderStatus: "to_ship",
		}
		// Update Order
		err = initializers.Db.Model(&order).Updates(&orderUpdate).Error
		if err != nil {
			log.Println("Failed to update db")
			return err
		}

		// Update Payment
		var paymentDone = true
		var paymentUpdate = models.PaymentUpdate{
			IsComplete: &paymentDone,
			PaymentAt:  &currentTime,
		}
		var payment models.Payment

		err = initializers.Db.Model(&payment).Where("stripe_session_id", stripeCheckoutSessionID).Updates(&paymentUpdate).Error
		if err != nil {
			log.Println("Failed to update db")
			return err
		}
	}
	return nil
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	GetAll[models.Payment](w, r, utils.EmptyScope)
}
