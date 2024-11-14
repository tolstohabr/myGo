package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mygo/internal/auth"
	"mygo/internal/model"
)

type Service interface {
	GetBanners() ([]model.Banner, error)
	CreateBanner(banner model.Banner) error
	UpdateBanner(banner model.Banner) error
	DeleteBanner(id int) error
}

// //////////////////////////////////////////////////////////////////
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

////////////////////////////////////////////////////////////////////

type Handler struct {
	service Service
}

// //////////////////////////////////////////////////////////////////////
func (h *Handler) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if creds.Username != "user" || creds.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

///////////////////////////////////////////////////////////////////////

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetBanners(c *gin.Context) {
	banners, err := h.service.GetBanners()
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

	if err := h.service.CreateBanner(banner); err != nil {
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

	if err := h.service.UpdateBanner(banner); err != nil {
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

	if err := h.service.DeleteBanner(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner deleted successfully"})
}
