package main

import (
	"egoist/app/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	
	router.Use(middleware.Logger)
	// register routes
	router.Route("/api/v1", func(r chi.Router) {
		routes.RegisterHealthRoutes(r)
		routes.RegisterAuthRoutes(r)
		routes.RegisterUserRoutes(r)
		routes.RegisterAzureRoutes(r)
		routes.RegisterAssetRoutes(r)
		routes.RegisterEntryRoutes(r)
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
