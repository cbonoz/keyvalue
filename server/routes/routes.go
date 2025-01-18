package routes

import (
	"keyvalue-api/email"
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
		apiKey := c.Request.Header.Get("x-api-key")
		app, err := s.queries.GetAPIKey(c.Request.Context(), apiKey)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("apiKey", app)
		c.Next()
	}
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	kv := r.Group("/api/kv")
	kv.Use(s.RequireValidKey())
	{
		kv.GET("/:key", s.getValue)
		kv.PUT("/:key", s.setValue)
		kv.DELETE("/:key", s.deleteValue)
	}

	// apiKeys := r.Group("/api/keys")
	// apiKeys.Use(middleware.RequireAuth())
	// {
	// 	apiKeys.POST("/", s.createAPIKey)
	// 	apiKeys.GET("/", s.listAPIKeys)
	// 	apiKeys.PATCH("/:id", s.updateAPIKey)
	// 	apiKeys.DELETE("/:id", s.deleteAPIKey)
	// }
}
