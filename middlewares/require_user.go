package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mooncorn/gshub-core/utils"
)

// RequireUser middleware ensures that the user's email is present in the context.
func RequireUser(c *gin.Context) {
	_, exists := c.Get("userEmail")
	if !exists {
		utils.HandleError(c, http.StatusUnauthorized, "Access unauthorized", errors.New("access restricted"), "null")
		c.Abort()
		return
	}

	c.Next()
}
