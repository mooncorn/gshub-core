package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mooncorn/gshub-core/utils"
)

// Ensures that the user's email is present in the context.
func RequireUser(c *gin.Context) {
	_, userEmailExists := c.Get("userEmail")
	_, userRoleExists := c.Get("userRole")

	userExists := userEmailExists && userRoleExists

	if !userExists {
		utils.HandleError(c, http.StatusUnauthorized, "Access unauthorized", errors.New("failed to require user: email or role is missing"), "null")
		c.Abort()
		return
	}

	c.Next()
}
