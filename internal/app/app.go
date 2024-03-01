package app

import (
	// "github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	// "github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	cols "github.com/xnpltn/hcc/internal/collections"
	"github.com/xnpltn/hcc/internal/handler"
	// tmp "github.com/xnpltn/hcc/internal/templates"
	// "github.com/xnpltn/hcc/internal/utils"
)

func App() *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		Save(app)
		e.Router.Static("/static", "assets")
		e.Router.GET("/", handler.Home(app))
		e.Router.GET("/services", handler.GetServices(app))
		e.Router.GET("/blog", handler.BlogGETHandler(app))
		e.Router.POST("/blog/create", handler.BlogPOSTHandler(app))
		e.Router.POST("/services/services", handler.CreateAService(app))
		e.Router.POST("/book", handler.BookAppointment(app))

		return nil
	})
	return app
}

var collections []*models.Collection = []*models.Collection{
	cols.AppointmentCollection,
	cols.BlogCollection,
	cols.MessageCollection,
	cols.QuoteCollection,
	cols.ServicesCollection,
}

func Save(app *pocketbase.PocketBase){
	for _, col := range collections {
		if err := app.Dao().SaveCollection(col); err != nil {
			continue
		}
	}
}
