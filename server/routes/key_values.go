package routes

import (
    "net/http"
    "keyvalue-api/models"
    "github.com/gin-gonic/gin"
)

func (s *Server) getValue(c *gin.Context) {
    userID := c.GetInt("user_id")
    appName := c.Param("app")
    key := c.Param("key")
    var keyValue models.KeyValue

    if err := s.db.First(&keyValue, "user_id = ? AND app_name = ? AND key = ?", userID, appName, key).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
        return
    }
    c.JSON(http.StatusOK, keyValue.Value)
}

func (s *Server) setValue(c *gin.Context) {
    userID := c.GetInt("user_id")
    appName := c.Param("app")
    key := c.Param("key")
    var req struct {
        Value string `json:"value"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    keyValue := models.KeyValue{
        UserID:  userID,
        AppName: appName,
        Key:     key,
        Value:   req.Value,
    }

    if err := s.db.Save(&keyValue).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, keyValue)
}

func (s *Server) deleteValue(c *gin.Context) {
    userID := c.GetInt("user_id")
    appName := c.Param("app")
    key := c.Param("key")

    result := s.db.Where("user_id = ? AND app_name = ? AND key = ?", userID, appName, key).Delete(&models.KeyValue{})
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
        return
    }
    c.Status(http.StatusNoContent)
}
