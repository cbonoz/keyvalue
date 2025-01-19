package middleware

import (
	"net/http"
	"strings"

	"keyvalue-api/setup"
	"keyvalue-api/sqlc_generated"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/gotrue-go/types"
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

func GetUser(
	c *gin.Context,
) *types.UserResponse {
	user, _ := c.Get("user")
	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}
	return user.(*types.UserResponse)
}

func GetApiKey(c *gin.Context) *sqlc_generated.ApiKey {
	apiKey, _ := c.Get("apiKey")
	if apiKey == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}
	return apiKey.(*sqlc_generated.ApiKey)
}
