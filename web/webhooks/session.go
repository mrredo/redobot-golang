package webhooks

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"main/config"
	"net/http"
	"os"
)

func SessionThing(c *gin.Context) {
	session1 := sessions.Default(c)
	yearlypremium := c.Query("pr")
	tok := session1.Get("token")
	if tok == nil {
		c.Redirect(http.StatusFound, "/auth/login?r=/checkout")
		return
	}
	amount := 2.99
	interval := "month"
	intervalcount := 1
	switch yearlypremium {
	case "year":
		yearlypremium = "prod_ObrYinRXqhcjOF"
		amount = 28.99
		interval = "year"
		break
	case "month":
		yearlypremium = "prod_ObrVDdHPA22IHJ"
	default:
		yearlypremium = "prod_ObrVDdHPA22IHJ"
	}
	user, err := config.AuthClient.GetUser(config.Sessions[tok.(string)])
	if err != nil {
		c.Redirect(http.StatusFound, "/auth/login?r=/checkout")
		return
	}
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"paypal",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String("eur"),
					Product:    stripe.String(yearlypremium),
					UnitAmount: stripe.Int64(int64(amount * 100)),
					Recurring: &stripe.CheckoutSessionLineItemPriceDataRecurringParams{
						Interval:      stripe.String(interval),
						IntervalCount: stripe.Int64(int64(intervalcount)),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Metadata:   map[string]string{"userid": user.ID.String()},
		Mode:       stripe.String("subscription"),
		SuccessURL: stripe.String(os.Getenv("BASE_URL") + "/guilds"),
		CancelURL:  stripe.String(os.Getenv("BASE_URL") + "/guilds"),
	}

	checkoutSession, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Checkout Session"})
		return
	}

	// Redirect the user to the Checkout Session URL
	c.Redirect(http.StatusSeeOther, checkoutSession.URL)
}
