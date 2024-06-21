package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mooncorn/gshub-core/models"
	"github.com/mooncorn/gshub-core/utils"
)

// Ensures that the user's role is correct.
func RequireRole(role models.UserRole) func(c *gin.Context) {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			utils.HandleError(c, http.StatusUnauthorized, "Access unauthorized", errors.New("failed to require user role: role is missing"), "null")
			c.Abort()
			return
		}

		if userRole != role {
			utils.HandleError(c, http.StatusUnauthorized, "Access unauthorized", errors.New("failed to require user role: role is invalid"), "null")
			c.Abort()
			return
		}

		c.Next()
	}
}
