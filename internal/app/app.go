package app

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/xnpltn/hcc/internal/utils"
)

func App() *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		LoadRoutes(e, app)
		return nil
	})

	app.OnBeforeApiError().Add(func(e *core.ApiErrorEvent) error {
		utils.WriteToLogs(e.HttpContext)
		utils.WriteToLogs(e.Error)
		return nil
	})
	return app
}




