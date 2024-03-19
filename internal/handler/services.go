package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	m "github.com/pocketbase/pocketbase/models"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/templates/pages"
	"github.com/xnpltn/hcc/internal/templates/views"
	"github.com/xnpltn/hcc/internal/utils"
)

func ServicesGet(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		services := []models.Service{}
		err := app.Dao().
			DB().
			Select("*").
			From("services").
			All(&services)
		if err != nil {
			utils.WriteToLogs(err)
		}
		return views.Marketplace(services).Render(c.Request().Context(), c.Response().Writer)
	}
}

func ServiceGetOneHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		service := models.Service{}
		reviews := []models.Review{}

		if err := app.Dao().DB().Select("*").From("services").Where(dbx.NewExp("id = {:id}", dbx.Params{"id": id})).One(&service); err != nil {
			
			utils.WriteToLogs(err)
		}
		if err := app.Dao().DB().Select("*").From("reviews").Where(dbx.NewExp("service_id = {:id}", dbx.Params{"id": id})).All(&reviews); err != nil {
			
			utils.WriteToLogs(err)
		}

		service.Reviews = reviews

		authCookie, err := c.Cookie("auth_token")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				utils.WriteToLogs(err)
			}
		}

		user := new(models.User)

		if authCookie != nil {
			authRecord, err := app.Dao().FindAuthRecordByToken(authCookie.Value, app.Settings().RecordAuthToken.Secret)
			if err != nil {
					
					utils.WriteToLogs(err)
					return pages.ServicePage(service, false, *user).Render(c.Request().Context(), c.Response().Writer)
			}
			user.Email = authRecord.Email()
			user.Username = authRecord.Username()
			user.Verified = authRecord.Verified()
			
			return pages.ServicePage(service, true, *user).Render(c.Request().Context(), c.Response().Writer)
			
		}
		return pages.ServicePage(service, false, *user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func PostReviewHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		values, err := c.FormValues()
		if err != nil {
			utils.ResolveErr(c, err)
		}
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			utils.ResolveErr(c, err)
		}
		userRecord, err := app.Dao().FindAuthRecordByToken(cookie.Value, app.Settings().RecordAuthToken.Secret)
		if err != nil {
			utils.ResolveErr(c, err)
		}
		rating, err := strconv.Atoi(values["rating"][0])
		if err != nil {
			utils.ResolveErr(c, err)
		}
		reviewCollection, err := app.Dao().FindCollectionByNameOrId("reviews")
		if err != nil {
			utils.ResolveErr(c, err)
		}

		record := m.NewRecord(reviewCollection)
		record.Set("user", userRecord.Get("name"))
		record.Set("review", values.Get("review"))
		record.Set("rating", rating)
		record.Set("service_id", values.Get("serviceID"))

		if err := app.Dao().Save(record); err != nil {
			utils.ResolveErr(c, err)
		}
		return c.HTML(201, "<p class='text-blue-700'>Review Saved Succefully</p>")
	}
}