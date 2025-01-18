package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getValue(c *gin.Context) {
	apiKey := s.RequireValidKey()
	// TODO:
	c.JSON(http.StatusOK, apiKey)
}

func (s *Server) setValue(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func (s *Server) deleteValue(c *gin.Context) {

	c.Status(http.StatusNoContent)
}
