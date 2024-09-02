package routes

import (
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterRevenueCatRoutes(r chi.Router){
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateWebhook)
		r.Post("/revenue-cat/webhook", controllers.RevenueCatWebhook)
	})	
}