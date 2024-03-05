package middleware

import (
	"BookingRoom/model/dto/json"
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {

	role := os.Getenv("ROLES")

	user, password, ok := c.Request.BasicAuth()
	if !ok {
		json.NewResponseUnauthorized(c, "Invalid Token", "00", "00")
		c.Abort()
		return
	}
	if user != role || password != role {
		json.NewResponseUnauthorized(c, "Unauthorized", "00", "00")
		c.Abort()
		return
	}
	c.Next()
}
