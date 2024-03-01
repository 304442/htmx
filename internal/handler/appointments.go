package handler

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	m "github.com/pocketbase/pocketbase/models"
)


func BookAppointment(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		user := c.FormValue("user")
		date := c.FormValue("date")
		timee := c.FormValue("time") 
		appointmentCollection, err := app.Dao().FindCollectionByNameOrId("appointment")
		if err != nil {
			return err
		}
		parsedTime, _:= time.Parse("2006-01-02 15:04" ,date+" "+timee)
		appointmentRecord := m.NewRecord(appointmentCollection)
		appointmentRecord.Set("user", user)
		appointmentRecord.Set("date", parsedTime.UTC())
		if err := app.Dao().SaveRecord(appointmentRecord); err != nil {
			fmt.Println(err)
		}
		return nil
	}
}