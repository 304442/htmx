package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/webhook"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/templates/pages"
	"github.com/xnpltn/hcc/internal/utils"
)

type Data struct {
	ProductID string `json:"_id" xml:"_id" form:"_id" query:"_id"`
}

type PaymentData struct {
	Description string `json:"description"`
}

func StripeonfigHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	conns := utils.GetAdminConfigs(app)
	return func(c echo.Context) error {
		if c.Request().Method != "GET" {
			c.String(400, "Invalid Method")
			return errors.New("invalid method")
		}
		return c.JSON(200, struct {
			PublishableKey string `json:"publishableKey"`
		}{
			PublishableKey: conns.StripePublishableKey,
		})
	}
}

func PaymentIntentHandler(app *pocketbase.PocketBase) echo.HandlerFunc {

	conns := utils.GetAdminConfigs(app)
	stripe.Key = conns.StripeSecretKey
	return func(c echo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, struct{ Error string }{Error: "error occured"})
		}
		data := new(Data)
		err = json.Unmarshal(b, data)
		if err != nil {
			utils.WriteToLogs(err)
			return err
		}
		service := models.Service{}
		err = app.Dao().DB().Select("*").From("services").Where(dbx.NewExp("id = {:id}", dbx.Params{"id": data.ProductID})).One(&service)
		if err != nil {
			utils.WriteToLogs(err)
			return c.JSON(404, struct{ Error string }{Error: "error occured"})
		}
		params := stripe.PaymentIntentParams{
			Amount:      stripe.Int64(int64(service.Price * 100)),
			Currency:    stripe.String(string(stripe.CurrencyUSD)),
			Description: stripe.String(fmt.Sprintf("Payment for %v", service.Title)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}

		pi, err := paymentintent.New(&params)
		if err != nil {
			if stripeErr, ok := err.(*stripe.Error); ok {
				log.Println("stripe error: ", stripeErr)
				utils.WriteToLogs(err)
				return c.JSON(404, struct{ Error string }{Error: "error occured"})
			} else {
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

func PaymentWebHook(app *pocketbase.PocketBase) echo.HandlerFunc {
	conns := utils.GetAdminConfigs(app)
	stripe.Key = conns.StripeSecretKey
	return func(c echo.Context) error {
		if c.Request().Method != "POST" {
			return c.String(404, "an error occured")
		}

		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(404, "an error occured")
		}
		event, err := webhook.ConstructEvent(b, c.Request().Header.Get("Stripe-Signature"), conns.StripeWebhookKey)
		if err != nil {
			utils.WriteToLogs(err)
			return c.String(400, "an error occured")
		}

		if event.Type == "payment_intent.succeeded" {
			conns := utils.GetAdminConfigs(app)
			data, err := event.Data.Raw.MarshalJSON()
			if err != nil {
				utils.WriteToLogs(err)
			}
			paymentData := new(PaymentData)
			if err := json.Unmarshal(data, paymentData); err != nil {
				utils.WriteToLogs(err)
			}
			// Notify the admin

			if err = utils.SendNotifications(app, "You Recieved a payement", fmt.Sprintf("Purchase: %s", paymentData.Description)); err != nil {
				fmt.Println(err)
			}
			if err := utils.SendMail(
				app,
				"admin",
				conns.AdminEmail,
				"You Recieved a payement",
				"<h1> Successfull Payment from a user</h1> <div> A product was succefuly purchaeds</div>",
			); err != nil {
				utils.WriteToLogs(err)
			}
		}
		return c.JSON(200, map[string]string{"success": "a successful payment"})
	}
}

// SuccessPage
func SuccessHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				utils.WriteToLogs(err)
			}
		}
		user := new(models.User)
		if cookie != nil {
			authRecord, err := app.Dao().FindAuthRecordByToken(cookie.Value, app.Settings().RecordAuthToken.Secret)
			if err != nil {
				utils.WriteToLogs(err)
				return pages.ProfilePage(models.User{}).Render(c.Request().Context(), c.Response().Writer)
			}
			user.Email = authRecord.Email()
			user.Username = authRecord.Username()
			user.Verified = authRecord.Verified()
			user.Name = authRecord.GetString("name")
		}
		utils.SendMail(app, user.Name, user.Email, "We Recived your payment", `
		<div >
			<h1>Successful payment</h1>
			<p>We have Revieved your payment</p>
		</div>
		`,
		)
		return pages.SuccessPage(*user).Render(c.Request().Context(), c.Response().Writer)
	}
}
