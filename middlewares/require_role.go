package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mooncorn/gshub-core/utils"
)

// Ensures that the user's role is correct.
func RequireRole(role string) func(c *gin.Context) {
	return func(c *gin.Context) {
		userRole := c.GetString("userRole")

		if !strings.EqualFold(userRole, role) {
			utils.HandleError(c, http.StatusUnauthorized, "Access unauthorized", errors.New("failed to require user role: role is invalid"), "null")
			c.Abort()
			return
		}

		c.Next()
	}
}
