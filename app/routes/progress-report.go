package routes

import (
	"egoist/app/controllers"
	"egoist/app/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterProgressReportRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthenticateJWT)
		r.Get("/progress-report/{reportId}", controllers.GetReport)
		// r.Get("/progress-report/all", controllers.GetReports) # not needed rn
	})
}