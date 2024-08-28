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

var MIMETYPES = map[string]string {
	"image/jpeg":  ".jpg",
    "image/pjpeg": ".jpg",  // Older MIME type for JPEG
    "image/png":   ".png",
    "image/gif":   ".gif",
    "image/bmp":   ".bmp",
    "image/webp":  ".webp",
    "image/tiff":  ".tiff",
    "image/x-tiff":".tiff",
    "image/vnd.microsoft.icon": ".ico", // Used for icons
    "image/x-icon": ".ico", // Alternative MIME type for icons
    "image/svg+xml": ".svg", // Scalable Vector Graphics
    "image/heic":   ".heic", // High Efficiency Image Coding (HEIC) for iOS
    "image/heif":   ".heif", // High Efficiency Image Format (HEIF)
    "image/avif":   ".avif", // AV1 Image File Format (used on Android and modern browsers)
}

func Map[T structs.Assets, V structs.Assets](array []T, fn func(index int, item T) V) []V{
	res := []V{}
	for index, asset := range array{
		res = append(res, fn(index, asset))
	}

	return res
}

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

func ReturnJson(w http.ResponseWriter, data any, status int){
		// return tokens to client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
	
		json.NewEncoder(w).Encode(data)
} 