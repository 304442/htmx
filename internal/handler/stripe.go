package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tmp "github.com/xnpltn/hcc/internal/templates"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/webhook"
	"github.com/xnpltn/hcc/internal/utils"
)



	



func StripeonfigHandler(app *pocketbase.PocketBase) echo.HandlerFunc{
	godotenv.Load()
	return func(c echo.Context) error {
		if c.Request().Method != "GET"{
			c.String(400, "Invalid Method")
			return errors.New("invalid method")
		}

		c.JSON(200, struct {
			PublishableKey string `json:"publishableKey"`
		}{
			PublishableKey: os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		})

		return nil
	}
}


func PaymentIntentHandler(app *pocketbase.PocketBase) echo.HandlerFunc{
	godotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return func(c echo.Context) error {
		params := stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1099),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}

		pi, err := paymentintent.New(&params)
		if err != nil{
			if stripeErr, ok := err.(*stripe.Error); ok{
				log.Println("stripe error: ", stripeErr)
				return c.JSON(404, struct{Error string}{Error: "error occured"})
			}else{
				return c.String(404, "error occured pi")
			}
		}


		return c.JSON(http.StatusOK, struct {
			ClientSecret string `json:"client_secret"`
		}{
			ClientSecret: pi.ClientSecret,
		})
	}
}

func PaymentWebHook(app *pocketbase.PocketBase) echo.HandlerFunc{
	godotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return func(c echo.Context) error {
		if c.Request().Method != "POST"{
			return c.String(404, "an error occured")
		}

		b, err := io.ReadAll(c.Request().Body)
		if err!= nil{
			return c.String(404, "an error occured")
		}

		event, err := webhook.ConstructEvent(b, c.Request().Header.Get("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	
		if err!= nil{
			return c.String(400, "an error occured")
		}
		if event.Type == "payment_intent.succeeded"{
			// add to db
			fmt.Println("Payment Succed")
		}
		return nil
	}
}


func SuccessHandler(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {	
		return utils.Render(tmp.Success(), c)
	}
}