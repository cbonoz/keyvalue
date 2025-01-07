package routes

import (
	"keyvalue-api/middleware"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db          *gorm.DB
	brevoClient *brevo.APIClient
}

func NewServer(db *gorm.DB, brevoClient *brevo.APIClient) *Server {
	return &Server{
		db: db,
		brevoClient: brevoClient,
	}
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	kv := r.Group("/api/kv")
	kv.Use(middleware.RequireAuth())
	{
		kv.GET("/:key", s.getValue)
		kv.PUT("/:key", s.setValue)
		kv.DELETE("/:key", s.deleteValue)
	}

	apiKeys := r.Group("/api/keys")
	apiKeys.Use(middleware.RequireAuth())
	{
		apiKeys.POST("/", s.createAPIKey)
		apiKeys.GET("/", s.listAPIKeys)
		apiKeys.PATCH("/:id", s.updateAPIKey)
		apiKeys.DELETE("/:id", s.deleteAPIKey)
	}
}
