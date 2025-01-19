package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"keyvalue-api/middleware"
	"keyvalue-api/sqlc_generated"
	"keyvalue-api/util"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/gin-gonic/gin"
)

type CreateAppRequest struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) createApp(c *gin.Context) {
	var req CreateAppRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUser(c)

	var createdByUuid pgtype.UUID
	createdByUuid.Scan(user.ID.String())

	app, err := s.queries.CreateApp(c, sqlc_generated.CreateAppParams{
		CreatedByUserID: createdByUuid,
		Name:            req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

func (s *Server) listApps(c *gin.Context) {
	user := middleware.GetUser(c)

	var createdBy pgtype.UUID
	createdBy.Scan(user.ID.String())

	apps, err := s.queries.ListUserApps(c, createdBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(apps) == 0 {
		// return empty list
		c.JSON(http.StatusOK, []sqlc_generated.App{})
	} else {
		c.JSON(http.StatusOK, apps)
	}
}

func (s *Server) getApp(c *gin.Context) {
	appId, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "app id must be an integer"})
		return
	}

	app, err := s.queries.GetApp(c, int32(appId))
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "app not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// delete app
func (s *Server) deleteApp(c *gin.Context) {
	appId := util.Int32(c.Param("id"))
	user := middleware.GetUser(c)

	params := sqlc_generated.DeleteAppParams{
		ID:              appId,
		CreatedByUserID: util.ConvertUuid(user.ID),
	}

	err := s.queries.DeleteApp(c, params)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "app not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// get app keys with obfuscated keys
func (s *Server) getAppKeys(c *gin.Context) {
	appId := util.Int32(c.Param("id"))

	keys, err := s.queries.ListKeyValues(c, appId)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "app not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, keys)
}
