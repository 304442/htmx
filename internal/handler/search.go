package handler

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/templates/views"
	"github.com/xnpltn/hcc/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SearchIn struct {
	SearchKeyWrd string `param:"search" query:"search" form:"search" json:"search" xml:"search"`
}

func SearchHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		services := []models.Service{}
		if err := app.Dao().DB().Select("*").From("services").All(&services); err != nil {
			utils.WriteToLogs(err)
		}
		search := new(SearchIn)
		c.Bind(search)
		searchRelated := []models.Service{}
		for _, service := range services {
			if search.SearchKeyWrd != "" {
				caser := cases.Title(language.English)
				if strings.Contains(service.Title, search.SearchKeyWrd) ||
					strings.Contains(service.Title, caser.String(search.SearchKeyWrd)) ||
					strings.Contains(service.Description, search.SearchKeyWrd) ||
					strings.Contains(service.Description, caser.String(search.SearchKeyWrd)) {
					searchRelated = append(searchRelated, service)
				}
			}
		}
		return views.Search(searchRelated).Render(c.Request().Context(), c.Response().Writer)
	}
}
