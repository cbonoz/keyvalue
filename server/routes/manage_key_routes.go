package routes

import (
	"crypto/rand"
	"encoding/base64"
	"keyvalue-api/sqlc_generated"
	"keyvalue-api/util"
	"net/http"

	"keyvalue-api/middleware"

	"github.com/gin-gonic/gin"
)

type CreateAPIKeyRequest struct {
	AppID int32 `json:"app_id" binding:"required"`
}

func generateAPIKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *Server) createAPIKey(c *gin.Context) {
	var req CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUser(c)

	// Verify app exists and user has access
	hasAccess, err := s.queries.ValidateAppOwnership(c, sqlc_generated.ValidateAppOwnershipParams{
		ID:              req.AppID,
		CreatedByUserID: util.ConvertUuid(user.ID),
	})

	if err != nil || !hasAccess {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	apiKey, err := generateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate key"})
		return
	}

	key, err := s.queries.CreateAPIKey(c, sqlc_generated.CreateAPIKeyParams{
		AppID: req.AppID,
		Key:   apiKey,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, key)
}

func (s *Server) getAPIKey(c *gin.Context) {
	keyStr := c.Param("key")
	key, err := s.queries.GetAPIKey(c, keyStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
		return
	}

	// Update last used timestamp
	if err := s.queries.UpdateAPIKeyLastUsed(c, key.ID); err != nil {
		// Log error but don't fail the request
		c.Error(err)
	}

	c.JSON(http.StatusOK, key)
}
