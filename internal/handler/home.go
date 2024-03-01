package handler

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/utils"
	tmp "github.com/xnpltn/hcc/internal/templates"
)

func Home(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		return utils.Render(tmp.Index(false), c)
	}
}