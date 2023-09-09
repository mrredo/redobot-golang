package webhooks

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"main/config"
	"net/http"
)

func SessionThing(c *gin.Context) {
	session1 := sessions.Default(c)
	tok := session1.Get("token")
	if tok == nil {
		c.Redirect(http.StatusFound, "/auth/login?r=/checkout")
		return
	}
	user, err := config.AuthClient.GetUser(config.Sessions[tok.(string)])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization is required"})
		return
	}
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"paypal",
			"alipay",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String("eur"),
					Product:    stripe.String("prod_ObWxPH6p9NyZ3Z"),
					UnitAmount: stripe.Int64(10_000_000), // Amount in cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Metadata:   map[string]string{"userid": user.ID.String()},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("https://your-website.com/success"),
		CancelURL:  stripe.String("https://your-website.com/cancel"),
	}

	checkoutSession, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Checkout Session"})
		return
	}

	// Redirect the user to the Checkout Session URL
	c.Redirect(http.StatusSeeOther, checkoutSession.URL)
}
