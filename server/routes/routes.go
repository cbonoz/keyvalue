package routes

import (
	"keyvalue-api/email"
	"keyvalue-api/middleware"
	"keyvalue-api/sqlc_generated"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries    *sqlc_generated.Queries
	appEmailer *email.AppEmailer
}

func NewServer(queries *sqlc_generated.Queries, brevoClient *brevo.APIClient) *Server {
	appEmailer := &email.AppEmailer{
		BrevoClient: brevoClient,
	}
	return &Server{
		queries:    queries,
		appEmailer: appEmailer,
	}
}

func (s *Server) RequireValidKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyString := c.Request.Header.Get("x-api-key")
		apiKey, err := s.queries.GetAPIKey(c.Request.Context(), apiKeyString)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("apiKey", apiKey)
		c.Next()
	}
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Key-Value routes with API key authentication
	kv := api.Group("/kv")
	kv.Use(s.RequireValidKey())
	{
		kv.GET("/:key", s.getValue)
		kv.PUT("/:key", s.setValue)
		kv.POST("/delete", s.deleteKeys)
	}

	// App and API Key management routes with user authentication
	authenticated := api.Group("")
	authenticated.Use(middleware.RequireAuth())
	{
		// App management
		apps := authenticated.Group("/apps")
		{
			apps.POST("", s.createApp)
			apps.GET("", s.listApps)
			apps.GET("/:id", s.getApp)
			apps.DELETE("/:id", s.deleteApp)
			apps.GET("/:id/keys", s.getAppKeys)
			apps.POST("/:id/keys", s.createAPIKey)
			apps.POST("/:id/keys/delete", s.deleteKeys)
		}
	}
}
