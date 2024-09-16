package main

import (
	"egoist/app"
	"egoist/app/routes"
	"egoist/internal/database"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)


func main() {
	global := app.NewGlobal(database.ConnectDB())

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/api/v1", func(r chi.Router) {
		routes.RegisterHealthRoutes(r)
		routes.RegisterAuthRoutes(r, global)
		routes.RegisterUserRoutes(r, global)
		routes.RegisterAWSRoutes(r, global)
		routes.RegisterAssetRoutes(r, global)
		routes.RegisterEntryRoutes(r, global)
		routes.RegisterRevenueCatRoutes(r, global)
	})

	var port string
	if port = os.Getenv("PORT"); port == "" {
		panic("The environment variable PORT wasn't provided.")
	}

	fmt.Printf("Listening on port %s....\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
