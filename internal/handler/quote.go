package handler

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	m "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/xnpltn/hcc/internal/utils"
)

func QuotePost(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {

		values, err := c.FormValues()
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusCreated, `
				<p class="text-red-700 text-center p-2 border border-gray-700">
					Something went wrong
				</p>
		`)
		}

		collection, err := app.Dao().FindCollectionByNameOrId("quotes")
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusCreated, `
				<p class="text-red-700 text-center p-2 border border-gray-700">
					Something went wrong
				</p>
			`)
		}

		record := m.NewRecord(collection)
		record.Set("name", values.Get("name"))
		record.Set("email", values.Get("email"))
		record.Set("phone", values.Get("phone"))
		record.Set("role", values.Get("role"))
		record.Set("companyname", values.Get("cname"))
		record.Set("companysize", values.Get("csize"))
		record.Set("industry", values.Get("industry"))
		record.Set("country", values.Get("country"))
		record.Set("usecase", values.Get("usecase"))
		record.Set("description", values.Get("description"))

		if err := app.Dao().SaveRecord(record); err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusCreated, `
			<p class="text-red-700 text-center p-2 border border-gray-700">
				Something went wrong
			</p>
			`)
		}

		// email to admin
		conns := utils.GetAdminConfigs(app)
		emailtoAdminMessage := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Name: "Admin",Address: conns.AdminEmail}},
			Subject: "A user requested a free quote",
			HTML:    fmt.Sprintf("<h1>A user requested a free quote</h1><div>A user with email %s requested a free quote, descriptions:  <p>%s</p> </div>", values.Get("email"), values.Get("description")),
		}
		
		if err = app.NewMailClient().Send(emailtoAdminMessage); err != nil {
			utils.WriteToLogs(err)
		}

		// email to user
		emailToUserMessage := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Name: values.Get("name"), Address: values.Get("email")}},
			Subject: "REQUEST RECEIVED.",
			HTML:    `<h1>Request received</h1><div>A request for a a free quete was succeccfully recieved, we will reach out as soon as possible </p> </div>`,
		}
		
		if err = app.NewMailClient().Send(emailToUserMessage); err != nil {
			utils.WriteToLogs(err)
		}

		return c.HTML(http.StatusCreated, `
		<p class="text-blue-700 text-center p-2 border border-gray-700">
			
		</p>
		`)
	}
}
