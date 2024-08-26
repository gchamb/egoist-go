package middlewares

import (
	"context"
	"egoist/internal/utils"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtToken = strings.Trim(strings.Split(jwtToken, "Bearer")[1], " ")
		if jwtToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
			
		userId, err := utils.VerifyToken(jwtToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	  })
}