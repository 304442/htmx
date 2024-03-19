package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/models"
	tmp "github.com/xnpltn/hcc/internal/templates"
	"github.com/xnpltn/hcc/internal/templates/pages"
	"github.com/xnpltn/hcc/internal/utils"
)

func Home(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		counties := []models.Country{}
		activeHours := []models.ActiveHours{}

		
		cookie, err := c.Cookie("auth_token")

		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				utils.WriteToLogs(err)
			}
		}
		err = app.
			Dao().
			DB().
			Select("*").
			From("activeHours").
			Where(dbx.NewExp("taken = {:taken}", dbx.Params{"taken": false})).
			OrderBy("active ASC").
			All(&activeHours)
		if err != nil {
			utils.WriteToLogs(err)
		}
		if err := app.
			Dao().
			DB().
			Select("*").
			From("countries").
			OrderBy("name ASC").
			All(&counties); err != nil {
			utils.WriteToLogs(err)
		}

		user := new(models.User)
		if cookie != nil {
			userRecord, err := app.Dao().FindAuthRecordByToken(cookie.Value, app.Settings().RecordAuthToken.Secret)
			if err != nil {
				utils.WriteToLogs(err)
				return tmp.Base(counties, false, activeHours, *user).Render(c.Request().Context(), c.Response().Writer)
			}
			user.Email = userRecord.Email()
			user.Username = userRecord.Username()
			user.Verified = userRecord.Verified()

			return tmp.Base(counties, true, activeHours, *user).Render(c.Request().Context(), c.Response().Writer)
		} else {
			return tmp.Base(counties, false, activeHours, *user).Render(c.Request().Context(), c.Response().Writer)
		}
	}
}

func ProfilePage(app *pocketbase.PocketBase) echo.HandlerFunc {
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
		return pages.ProfilePage(*user).Render(c.Request().Context(), c.Response().Writer)
	}
}
