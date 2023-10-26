package webhooks

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"io/ioutil"
	"main/structs"
	"net/http"
	"os"
)

func HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse webhook body json"})
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing webhook JSON"})
			return
		}
		//userEmail := paymentIntent.Customer.Email
		//productID := paymentIntent.ID

		paymentIntentmap := map[string]any{}
		json.Unmarshal(event.Data.Raw, &paymentIntentmap)
	// Then define and call a func to handle the successful payment intent.
	// handlePaymentIntentSucceeded(paymentIntent)
	case "checkout.session.completed":
		paymentIntentmap := map[string]any{}
		json.Unmarshal(event.Data.Raw, &paymentIntentmap)
		id, ok := paymentIntentmap["metadata"].(map[string]any)["userid"]
		if !ok {
			fmt.Println(0)
			return
		}
		//cus, ok1 := paymentIntentmap["customer_id"]
		//if !ok1 {
		//	fmt.Println(paymentIntentmap)
		//	return
		//}
		sub, ok2 := paymentIntentmap["subscription"]
		if !ok2 {
			fmt.Println(2)
			return
		}
		expat, ok3 := paymentIntentmap["expires_at"]
		if !ok3 {
			fmt.Println(3)
			return
		}
		fmt.Println(paymentIntentmap)
		prem := structs.NewPremium(id.(string), sub.(string), expat.(float64))
		prem.RegisterNewCustomer()

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing webhook JSON"})
			return
		}
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handlePaymentMethodAttached(paymentMethod)
	// ... handle other event types
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed successfully"})
}
