package app

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/xnpltn/hcc/internal/handler"
)

func LoadRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.Static("/static", "assets")

	// views routes(Fronted Routes)
	loadViewsRoutes(e, app)

	// API routes(Backed Routes)
	loadAPIRoutes(e, app)

	// stripe routes
	loadStripeRoutes(e, app)

}

func loadViewsRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.GET("/", handler.Home(app))
	e.Router.GET("/marketplace", handler.ServicesGet(app))
	e.Router.GET("/blog", handler.BlogsGETHandler(app))
	e.Router.GET("/service/:id", handler.ServiceGetOneHandler(app))
	e.Router.GET("/blog/:id", handler.BlogGETOneHander(app))
	e.Router.GET("/faqs", handler.FaqGETHandler(app))
	e.Router.GET("/auth/login", handler.LoginForn())
	e.Router.GET("/auth/signup", handler.SignupForn())
	e.Router.GET("account/profile", handler.ProfilePage(app))
	e.Router.GET("/auth/verify", handler.VerifyRequestHandler(app))
	e.Router.GET("/auth/verify/:token", handler.VerifyEmail(app))
	e.Router.GET("/auth/cofirm-pass-rest/:token", handler.ChangePasshandler(app))
}

func loadAPIRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.POST("/appointment", handler.BookAppointment(app))
	e.Router.POST("/faqs", handler.FaqPostSpecific(app))
	e.Router.POST("/quote", handler.QuotePost(app))
	e.Router.POST("/signup", handler.SignupHandler(app))
	e.Router.POST("/login", handler.LoginHandler(app))
	e.Router.POST("/search", handler.SearchHandler(app))
	e.Router.POST("/review", handler.PostReviewHandler(app))
	e.Router.POST("/update", handler.UpdateProfile(app))
	e.Router.GET("/logout", handler.HandlerLogout(app))
	e.Router.POST("/passreset", handler.PassRequestHandler(app))
	e.Router.POST("/reset-pass", handler.HandleResetpassword(app))
}

func loadStripeRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	e.Router.GET("/config", handler.StripeonfigHandler(app))
	e.Router.POST("/cpi", handler.PaymentIntentHandler(app))
	e.Router.POST("/webhook", handler.PaymentWebHook(app))
	e.Router.GET("/success", handler.SuccessHandler(app))
}
