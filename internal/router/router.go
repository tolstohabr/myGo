package router

import (
	"mygo/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetBanners(c *gin.Context)
	CreateBanner(c *gin.Context)
	UpdateBanner(c *gin.Context)
	DeleteBanner(c *gin.Context)
	///////////////////////////////////////////////////////////////////////////////////
	Login(c *gin.Context)
	///////////////////////////////////////////////////////////////////////////////////
}

// ////////////////////////////////////////////////////////////////////////////////////
type YourHandler struct {
	// можно добавить необходимые зависимости
}

//////////////////////////////////////////////////////////////////////////////////////

type HttpRouter struct {
	router *gin.Engine
}

func NewHttpRouter() *HttpRouter {
	router := gin.Default()

	//после комментирования логгера в миделвере это не нужно
	//router.Use(middleware.LoggerMiddleware())

	return &HttpRouter{router: router}
}

func (r *HttpRouter) Register(handler Handler) {
	///////////////////////////////////////////////////////////////////////////////////
	r.router.POST("/login", handler.Login)

	protected := r.router.Group("/")
	protected.Use(middleware.JWTMiddleware)

	protected.GET("/banners", handler.GetBanners)
	protected.POST("/banners/create", handler.CreateBanner)
	protected.PUT("/banners/update", handler.UpdateBanner)
	protected.DELETE("/banners/delete", handler.DeleteBanner)
	//////////////////////////////////////////////////////////////////////////////////
}

func (r *HttpRouter) Run(address string) error {
	return r.router.Run(address)
}

/*
// по идее это дублирование потому что в хендлере есть это
// /////////////////////////////////////////////////////////////////////////////////////
func (h *YourHandler) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if credentials.Username != "user" || credentials.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

///////////////////////////////////////////////////////////////////////////////////////
*/
