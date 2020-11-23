package innkeeper

import (
	gin "github.com/gin-gonic/gin"

	// "log"
	"net/http"
)

// GetTestValidation is ...
func GetTestValidation(c *gin.Context) {
	c.String(http.StatusOK, "User is valid.")
}
