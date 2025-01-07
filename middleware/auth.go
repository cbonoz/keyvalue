package middleware

import (
	"net/http"
	"strings"

	"keyvalue-api/setup"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)

		client := setup.SupabaseClient.Auth.WithToken(token)
		user, err := client.GetUser()

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
