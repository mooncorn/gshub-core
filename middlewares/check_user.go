package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Checks for the authorization header and attaches the user's email to the context.
func CheckUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Next()
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		c.Next()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.Next()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if role, exists := claims["id"].(string); exists {
			c.Set("userID", role)
		}
		if email, exists := claims["email"].(string); exists {
			c.Set("userEmail", email)
		}
		if role, exists := claims["role"].(string); exists {
			c.Set("userRole", role)
		}
	}

	c.Next()
}
