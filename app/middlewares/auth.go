package middlewares

import (
	"context"
	"egoist/internal/utils"
	"net/http"
	"os"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.ValidateJWT(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		ctx := context.WithValue(r.Context(), "uid", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	  })
}

func AuthenticateWebhook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rcAuthToken := r.Header.Get("Authorization")
		if rcAuthToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rcAuthToken = strings.Trim(strings.Split(rcAuthToken, "Bearer")[1], " ")
		if rcAuthToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
			
		if rcAuthToken != os.Getenv("REVENUE_CAT_WEBHOOK_KEY"){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	  })
}