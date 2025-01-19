package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"keyvalue-api/constants"
	"keyvalue-api/sqlc_generated"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
)

func parseJWTToken(token string, hmacSecret []byte) (Claims jwt.MapClaims, err error) {
	// Parse the token and validate the signature
	t, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
		return *claims, nil
	} else {
		err := fmt.Errorf("invalid token: %v", err)
		return nil, err
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Default().Println("No authorization header provided")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		// authenticate jwt
		hmac := []byte(constants.SUPABASE_JWT_SECRET)
		token, err := parseJWTToken(tokenString, hmac)
		if err != nil {
			log.Default().Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Convert claims to map string
		claims := make(map[string]interface{})
		for k, v := range token {
			claims[k] = v
		}

		var user types.User
		claimString, err := json.Marshal(claims)
		if err != nil {
			log.Default().Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Get user from claims
		err = json.Unmarshal(claimString, &user)

		if err != nil {
			log.Default().Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user.ID, err = uuid.Parse(claims["sub"].(string))
		if err != nil {
			log.Default().Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func GetUser(
	c *gin.Context,
) types.User {
	user, _ := c.Get("user")
	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return types.User{}
	}
	return user.(types.User)
}

func GetApiKey(c *gin.Context) *sqlc_generated.ApiKey {
	apiKey, _ := c.Get("apiKey")
	if apiKey == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}
	return apiKey.(*sqlc_generated.ApiKey)
}
