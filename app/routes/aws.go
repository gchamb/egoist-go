package routes

import (
	"egoist/app"
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterAWSRoutes(r chi.Router, global *app.Globals) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Get("/aws/upload", controllers.GenerateUploadPresignedUrl(global))
	})
}
