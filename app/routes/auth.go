package routes

import (
	"egoist/app"
	"egoist/app/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, global *app.Globals) {
	r.Post("/auth/google", controllers.SignInWithGoogle(global))
	r.Post("/auth/apple", controllers.SignInWithApple(global))
	r.Post("/auth/signin", controllers.SignInWithEmail(global))
	r.Post("/auth/signup", controllers.SignUpWithEmail(global))
}
