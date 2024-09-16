package routes

import (
	"egoist/app"
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterEntryRoutes(r chi.Router, global *app.Globals) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Put("/entry", controllers.PutEntry(global))
	})
}
