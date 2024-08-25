package utils

import (
	"egoist/internal/structs"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.Claims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateFreshToken(claims jwt.Claims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateTokens(id string) (structs.AuthResponse, error) {
	jwtClaims := jwt.RegisteredClaims{Subject: id, ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0,0, 7))}
	jwtToken, err := GenerateJWT(jwtClaims)
	if err != nil {
		return structs.AuthResponse{}, err
	}

	refreshTokenClaims := jwt.RegisteredClaims{Subject: id, ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0,0,14))}
	refreshToken, err := GenerateJWT(refreshTokenClaims)
	if err != nil {
		return structs.AuthResponse{}, err
	}
	return structs.AuthResponse{
		JwtToken: jwtToken,
		FreshToken: refreshToken,
	}, err
}

func VerifyToken(tokenString string) (string, error){
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	 })
	
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return token.Claims.GetSubject()	
}

func ReturnJson(w http.ResponseWriter, data any){
		// return tokens to client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(data)
} 