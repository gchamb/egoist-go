package main

import (
	"fmt"
	"net/http"
	"os"

	"egoist/app/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	
	// register routes
	router.Route("/api/v1", func(r chi.Router) {
		routes.RegisterHealthRoutes(r)
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
