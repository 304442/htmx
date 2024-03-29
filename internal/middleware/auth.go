package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)



func AuthMiddleware(app *pocketbase.PocketBase) echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return nil
		}
	}
}