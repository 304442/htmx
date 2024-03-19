package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	m "github.com/pocketbase/pocketbase/models"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/utils"
)

func BookAppointment(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		values, err := c.FormValues()
		activeDate := models.ActiveHours{}
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong</p>")
		}

		appointmentCollection, err := app.Dao().FindCollectionByNameOrId("appointment")
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong</p>")
		}

		parsedTime, err := time.Parse("2006-01-02 15:04", values.Get("date"))
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong</p>")
		}
		appointmentRecord := m.NewRecord(appointmentCollection)
		appointmentRecord.Set("email", values.Get("email"))
		appointmentRecord.Set("name", values.Get("name"))
		appointmentRecord.Set("date", parsedTime.UTC())
		appointmentRecord.Set("free", values.Get("free") == "1")

		if err := app.
			Dao().
			DB().
			Select("*").
			From("activeHours").
			Where(
				dbx.NewExp("active = {:active}", dbx.Params{"active": fmt.Sprintf("%v:00.000Z", values.Get("date"))}),
			).
			One(&activeDate); err != nil {
			utils.WriteToLogs(err)
		}

		record, err := app.Dao().FindRecordById("activeHours", activeDate.ID)
		if err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong, Appointment not sent</p>")
		}

		record.Set("taken", true)
		record.Set("email", values.Get("email"))
		record.Set("name", values.Get("name"))

		if err := app.Dao().Save(record); err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong, Appointment not sent</p>")
		}

		if err := app.Dao().SaveRecord(appointmentRecord); err != nil {
			utils.WriteToLogs(err)
			return c.HTML(http.StatusOK, "<p>Something went wrong, Appointment not sent</p>")
		}

		// email to admin
		conns := utils.GetAdminConfigs(app)
		if err := utils.SendMail(
			app,
			"Admin",
			conns.AdminEmail,
			"A user booked an Appointment",
			fmt.Sprintf("<h1> Appointment Booked succeccfully</h1> <div> user with email %s booked and Appointment at %v </div>",values.Get("email"), parsedTime),
		); err != nil {
			utils.WriteToLogs(err)
		}

		// email to user
		if err := utils.SendMail(
			app,
			values.Get("name"),
			values.Get("email"),
			"Your Appointment was successully recieved",
			fmt.Sprintf("<h1> Appointment Booked succeccfully</h1> <div> You booked an Appointment at %v </div>", parsedTime),
		); err != nil {
			utils.WriteToLogs(err)
		}
		
		return c.HTML(http.StatusCreated, `
		<div class="text-center text-blue-700 p-2 border border-gray-700 rounded-md">
			<p>Request Sent successfully<br/>
				We will reach out soon
			</p>
		</div>
		`)
	}
}
