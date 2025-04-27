// internal/utils/jwt.go
package utils

import (
	"fmt"
	"habit-tracker/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(config.GetEnv("JWT_KEY", "")) // Change this to something secure!

// GenerateJWT creates a new JWT for the given user
func GenerateJWT(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expire in 24 hours
		"iss":      "Habit-tracker",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}

// ValidateJWT validates the provided JWT and returns the claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not extract claims")
	}

	return claims, nil
}
