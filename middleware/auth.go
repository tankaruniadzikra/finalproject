package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Token is missing"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["sub"].(float64))
		log.Printf("User ID from token: %d", userID)

		c.Set("user", userID)

		return next(c)
	}
}
