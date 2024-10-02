package routes

import (
	"egoist/app"
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, global *app.Globals) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Get("/user", controllers.GetUser(global))
		r.Delete("/user", controllers.DeleteUser(global))
		r.Patch("/user/onboard", controllers.OnboardUser(global))
		r.Patch("/user/update", controllers.UpdateUser(global))
	})
}
