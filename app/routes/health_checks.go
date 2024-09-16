package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterHealthRoutes(r chi.Router) {
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Up and running 🎉"))
	})
}
