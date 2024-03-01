package handler

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/xnpltn/hcc/internal/models"
	tmp "github.com/xnpltn/hcc/internal/templates"
	"github.com/xnpltn/hcc/internal/utils"
	m "github.com/pocketbase/pocketbase/models"
)




func BlogGETHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		blogs := []models.Blog{}
		err := app.Dao().DB().Select("title", "author", "content", "created").From("blog").Limit(10).All(&blogs)
		if err!= nil{
			fmt.Println(err)
		}
		return utils.Render(tmp.Blog(blogs), c)
		
	}
}


func BlogPOSTHandler(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {

		title :=c.FormValue("title")
		content := c.FormValue("content")
		author := c.FormValue("author") 
		collection, err := app.Dao().FindCollectionByNameOrId("blog")
		if err != nil {
			return err
		}

		record := m.NewRecord(collection)
		record.Set("title", title)
		record.Set("author", author)
		record.Set("content", content)

		if err := app.Dao().SaveRecord(record); err != nil {
			fmt.Println(err)
		}
		return utils.Render(tmp.Blog([]models.Blog{}), c)
	}
}

