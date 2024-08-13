package routes

import (
	"egoist/app/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router) {
	r.Post("/auth/google", controllers.SignInWithGoogle)
	r.Post("/auth/signin", controllers.SignInWithEmail )
	r.Post("/auth/signup", controllers.SignUpWithEmail)
}