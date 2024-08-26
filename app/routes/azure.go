package routes

import (
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterAzureRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Get("/azure/upload", controllers.GenerateUploadSaSUrl)
	})
}