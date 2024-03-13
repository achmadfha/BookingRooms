package middleware

import (
	"BookingRoom/model/dto/json"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(roles ...string) gin.HandlerFunc {
	secret := os.Getenv("SECRET_TOKEN")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Invalid Token", "00", "01")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			json.NewResponseUnauthorized(c, "Invalid Token", "00", "02")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			json.NewResponseUnauthorized(c, "Invalid Token", "00", "03")
			c.Abort()
			return
		}

		expirationTime := int64(claims["expired"].(float64))
		if time.Now().Unix() > expirationTime {
			json.NewResponseUnauthorized(c, "Token Expired", "00", "05")
			c.Abort()
			return
		}

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims["role"] {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseForbidden(c, "Forbidden", "00", "06")
			c.Abort()
			return
		}

		c.Next()
	}
}
