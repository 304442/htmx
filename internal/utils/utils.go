package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func RenderView(c echo.Context, view templ.Component, layoutPath string)error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response().Writer)
	}
	return nil
}


func ResolveErr(c echo.Context, err error) error{
	WriteToLogs(err)
	return c.HTML(200, "<p>Something went wrong, try again later</p>")
}



