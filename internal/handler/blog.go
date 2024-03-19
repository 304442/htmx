package handler

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/models"
	"github.com/xnpltn/hcc/internal/templates/pages"
	views "github.com/xnpltn/hcc/internal/templates/views"
	"github.com/xnpltn/hcc/internal/utils"
)

func BlogsGETHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		blogs := []models.Blog{}
		err := app.Dao().
			DB().
			Select("*").
			From("blog").
			OrderBy("created DESC").
			All(&blogs)
		if err != nil {
			utils.WriteToLogs(err)
			return views.Blog([]models.Blog{}).Render(c.Request().Context(), c.Response().Writer)
		}

		return views.Blog(blogs).Render(c.Request().Context(), c.Response().Writer)
	}
}

func BlogGETOneHander(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		blog := models.Blog{}

		if err := app.Dao().
			DB().
			Select("*").
			From("blog").
			Where(dbx.NewExp("id = {:id}", dbx.Params{"id": id})).
			One(&blog); err != nil {
			utils.WriteToLogs(err)
		}
		return pages.BlogPage(blog).Render(c.Request().Context(), c.Response().Writer)
	}
}
