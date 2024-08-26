package routes

import (
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Patch("/user/onboard", controllers.OnboardUser)
	})
}