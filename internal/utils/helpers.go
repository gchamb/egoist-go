package utils

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"egoist/internal/structs"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"encoding/pem"

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

func Map[T structs.Assets, V any](array []T, fn func(index int, item T) V) []V{
	res := []V{}
	for index, asset := range array{
		res = append(res, fn(index, asset))
	}

	return res
}

func parseApplePrivateKey() (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(os.Getenv("APPLE_JWT_SECRET")))
	// Parse the EC private key
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
				return nil, err
	}
	ecPrivateKey, _ := privateKey.(*ecdsa.PrivateKey)
	return ecPrivateKey, nil
}

func GenerateJWT(claims jwt.Claims, isApple bool) (string, error){
	var token *jwt.Token
	
	if isApple {
			token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
			key, err := parseApplePrivateKey()
			if err != nil {
				return "", err
			}
			return token.SignedString(key)
		}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateFreshToken(claims jwt.Claims) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateTokens(id string) (structs.AuthResponse, error) {
	jwtClaims := jwt.RegisteredClaims{Subject: id, ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0,0, 7))}
	jwtToken, err := GenerateJWT(jwtClaims, false)
	if err != nil {
		return structs.AuthResponse{}, err
	}

	refreshTokenClaims := jwt.RegisteredClaims{Subject: id, ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0,0,14))}
	refreshToken, err := GenerateJWT(refreshTokenClaims, false)
	if err != nil {
		return structs.AuthResponse{}, err
	}
	return structs.AuthResponse{
		JwtToken: jwtToken,
		FreshToken: refreshToken,
	}, err
}

func VerifyToken(tokenString string, isApple bool) (jwt.Claims, error){

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if isApple {
			kid := token.Header["kid"].(string)
			if kid == "" {
				return nil, errors.New("kid doesn't exist in header")
			}

			publicKey, err := FetchAppleKeysByKid(kid)
			fmt.Println(publicKey, "public key")
			if err != nil {
				return nil, err
			}

			return convertJWKToRSAPublicKey(publicKey)
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	 })
	
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims, nil
}

func FetchAppleKeysByKid(kid string) (structs.JWK,  error) {
	var jwks struct {
		Keys []structs.JWK `json:"keys"`
	}
	
	jwkURL := "https://appleid.apple.com/auth/keys"
	resp, err := http.Get(jwkURL)
	if err != nil {
		return jwks.Keys[0], err
	}
	defer resp.Body.Close()


	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return jwks.Keys[0], err
	}

	// search for kid
	for i := 0; i<len(jwks.Keys); i++ {
		if jwks.Keys[i].Kid == kid {
			return jwks.Keys[i], nil
		}
	}

	return jwks.Keys[0], errors.New("unable to find kid")
}

func convertJWKToRSAPublicKey(jwk structs.JWK) (*rsa.PublicKey, error) {
	// Decode the modulus (n) and exponent (e) from base64
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus (n): %w", err)
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent (e): %w", err)
	}

	// Convert the exponent to an integer
	e := big.NewInt(0).SetBytes(eBytes).Int64()

	// Create the rsa.PublicKey from modulus and exponent
	pubKey := &rsa.PublicKey{
		N: big.NewInt(0).SetBytes(nBytes),
		E: int(e),
	}

	return pubKey, nil
}

func ReturnJson(w http.ResponseWriter, data any, status int){
		// return tokens to client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
	
		json.NewEncoder(w).Encode(data)
} 