package routes

import (
	"egoist/app"
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterRevenueCatRoutes(r chi.Router, global *app.Globals) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateWebhook)
		r.Post("/revenue-cat/webhook", controllers.RevenueCatWebhook(global))
	})
}
