package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your-secret-key") // Use a secure key in production

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}