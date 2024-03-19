package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/mails"
	"github.com/xnpltn/hcc/internal/templates/pages"
	"github.com/xnpltn/hcc/internal/utils"
)



func VerifyEmail(app * pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		authRecord, err := app.Dao().FindAuthRecordByToken(c.PathParam("token"), app.Settings().RecordVerificationToken.Secret)
		if err!= nil{
			utils.WriteToLogs(err)
		}	
		if err := authRecord.SetVerified(true); err != nil{
				utils.WriteToLogs(err)	
		}
		if err := app.Dao().Save(authRecord); err != nil{
			utils.WriteToLogs(err)
		}
		return pages.VerifyEmail(authRecord.Email()).Render(c.Request().Context(), c.Response().Writer)
	}
}

func VerifyRequestHandler(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				utils.WriteToLogs(err)
			}
		}
		if cookie != nil {
			authRecord, err := app.Dao().FindAuthRecordByToken(cookie.Value, app.Settings().RecordAuthToken.Secret)
			if err != nil {
				utils.WriteToLogs(err)
				return c.HTML(200, "<p>Profile Updated Something went wrong</p>")
			}
			err = mails.SendRecordVerification(app, authRecord)
			if err != nil {
				utils.WriteToLogs(err)
				return c.HTML(200, "<p>Profile Updated Something went wrong</p>")
			}
		}
		return c.HTML(200, "<p class='text-center text-blue-700 my-2' >Email verification link sent.<br/> Check inbokx</p>")
	}
}