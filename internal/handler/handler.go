package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mygo/internal/model"
)

type Repository interface {
	GetBanners() ([]model.Banner, error)
	CreateBanner(banner model.Banner) error
	UpdateBanner(banner model.Banner) error
	DeleteBanner(id int) error
}

type Handler struct {
	repo Repository
}

func NewHander(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetBanners(c *gin.Context) {
	banners, err := h.repo.GetBanners()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, banners)
}

func (h *Handler) CreateBanner(c *gin.Context) {
	var banner model.Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateBanner(banner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Banner created successfully"})
}

func (h *Handler) UpdateBanner(c *gin.Context) {
	var banner model.Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateBanner(banner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner updated successfully"})
}

func (h *Handler) DeleteBanner(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	if err := h.repo.DeleteBanner(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner deleted successfully"})
}
