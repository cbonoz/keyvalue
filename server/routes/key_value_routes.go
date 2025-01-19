package routes

import (
	"keyvalue-api/middleware"
	"keyvalue-api/models"
	"keyvalue-api/sqlc_generated"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getValue(c *gin.Context) {
	apiKey := middleware.GetApiKey(c)

	value, err := s.queries.GetKeyValue(c.Request.Context(),
		sqlc_generated.GetKeyValueParams{
			AppID: apiKey.AppID,
			Key:   c.Param("key"),
		})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, value)
}

func (s *Server) setValue(c *gin.Context) {
	apiKey := middleware.GetApiKey(c)

	var body models.SetKeyValue

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertValue := sqlc_generated.UpsertKeyValueParams{
		AppID: apiKey.ID,
		Key:   body.Key,
		Value: body.Value,
	}

	// TODO: check space on account

	value, err := s.queries.UpsertKeyValue(c.Request.Context(), upsertValue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, value)
}

func (s *Server) deleteKeys(c *gin.Context) {
	apiKey := middleware.GetApiKey(c)

	var body models.DeleteKeys

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteParams := sqlc_generated.DeleteKeyValuesParams{
		AppID: apiKey.ID,
		Column2:  body.Keys,
	}

	err := s.queries.DeleteKeyValues(c.Request.Context(), deleteParams)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
