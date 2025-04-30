// internal/middleware/jwt_middleware.go
package middleware

import (
	"habit-tracker/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// JWTMiddleware verifies the JWT token passed in the Authorization header
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the token from the "Authorization" header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header is missing"})
		}

		// Split "Bearer <token>"
		tokenString := strings.Split(authHeader, " ")[1]

		// Validate the token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
		}

		// Store claims in context (for future use, e.g., in handler)
		c.Set("user", claims["username"])

		return next(c)
	}
}
