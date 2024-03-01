package handler

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/models"
	tmp "github.com/xnpltn/hcc/internal/templates"
	"github.com/xnpltn/hcc/internal/utils"
)


func GetServices(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		services := []models.Service{}
		
		err := app.Dao().DB().Select("*").From("services").All(&services)
		if err!= nil{
			fmt.Println(err)
		}
		return utils.Render(tmp.Services(services), c)
	}
}


func CreateAService(app *pocketbase.PocketBase) echo.HandlerFunc{
	return func(c echo.Context) error {
		return nil
	}
}