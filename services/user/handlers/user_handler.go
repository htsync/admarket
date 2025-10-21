package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/you/connect-market/services/user/models"
	"github.com/you/connect-market/services/user/repository"
)

type UserHandler struct {
	repo *repository.ProfileRepository
}

func NewUserHandler(repo *repository.ProfileRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	profile, err := h.repo.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req models.Profile
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = id
	if err := h.repo.Update(context.Background(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *UserHandler) SearchByTag(c *gin.Context) {
	tag := c.Query("tag")
	res, _ := h.repo.SearchByTag(context.Background(), tag)
	c.JSON(http.StatusOK, res)
}
