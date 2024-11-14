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
