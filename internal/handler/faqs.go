package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	m "github.com/pocketbase/pocketbase/models"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/templates/views"
	"github.com/xnpltn/hcc/internal/utils"
)

func FaqGETHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		faqs := []models.FAQ{}
		err := app.Dao().DB().Select("*").From("faqs").All(&faqs)
		if err != nil {
			utils.WriteToLogs(err)
		}
		return views.FaqView(faqs).Render(c.Request().Context(), c.Response().Writer)
	}
}

func FaqPostSpecific(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		collection, err := app.Dao().FindCollectionByNameOrId("questions")
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusInternalServerError, `
			<p style="text-align: center; 
				color: red; 
				border-width: 1px; 
				padding: 0.25rem;
				--tw-border-opacity: 1;
				border-color: rgb(31 41 55 / var(--tw-border-opacity));
				border-radius: 0.375rem;
				">
				An error occured don't worry, Not your fault.
			</p>
			`)
		}
		record := m.NewRecord(collection)
		record.Set("user", c.FormValue("email"))
		record.Set("message", c.FormValue("message"))
		if err := app.Dao().SaveRecord(record); err != nil {
			utils.ResolveErr(c, err)
		}

		// email to admin

		if err := utils.SendMail(
			app,
			"Admin",
			os.Getenv("ADMIN_EMAIL"),
			"A user asked a question",
			fmt.Sprintf("<h1>A user Asked a Question</h1><div>A user with email %s asked a question %s </div>", c.FormValue("email"), c.FormValue("message")),
		); err != nil {
			utils.WriteToLogs(err)
		}
		// email to user
		if err := utils.SendMail(
			app,
			"Admin",
			c.FormValue("email"),
			"Your Question Was Recieved",
			"<h1>Recieved</h1><div> Your Question was recieved and we will reach out to you as soon as possible.</div>",
		); err != nil {
			utils.WriteToLogs(err)
		}
		return c.HTML(http.StatusCreated, `
		<div class="flex items-center flex-col text-center">
		<p>Message Recieved Successfully <br/>
			We will get back to you soon
		</p>
		</div>
	`)
	}
}
