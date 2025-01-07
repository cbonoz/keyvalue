package routes

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"keyvalue-api/models"

	"github.com/gin-gonic/gin"
)

type CreateAPIKeyRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAPIKeyRequest struct {
	Name     *string `json:"name"`
	IsActive *bool   `json:"is_active"`
}

func generateAPIKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (s *Server) createAPIKey(c *gin.Context) {
	userID := c.GetInt("user_id")
	var req CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiKey := models.APIKey{
		UserID:   userID,
		Key:      generateAPIKey(),
		Name:     req.Name,
		IsActive: true,
	}

	if err := s.db.Create(&apiKey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, apiKey)
}

func (s *Server) listAPIKeys(c *gin.Context) {
	userID := c.GetInt("user_id")
	var keys []models.APIKey

	if err := s.db.Where("user_id = ?", userID).Find(&keys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (s *Server) updateAPIKey(c *gin.Context) {
	userID := c.GetInt("user_id")
	keyID := c.Param("id")
	var req UpdateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	result := s.db.Model(&models.APIKey{}).
		Where("user_id = ? AND id = ?", userID, keyID).
		Updates(updates)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "API key not found"})
		return
	}
	c.Status(http.StatusOK)
}

func (s *Server) deleteAPIKey(c *gin.Context) {
	userID := c.GetInt("user_id")
	keyID := c.Param("id")

	result := s.db.Where("user_id = ? AND id = ?", userID, keyID).
		Delete(&models.APIKey{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "API key not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
