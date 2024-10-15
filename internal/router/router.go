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
}

type HttpRouter struct {
	router *gin.Engine
}

func NewHttpRouter() *HttpRouter {
	router := gin.Default()

	router.Use(middleware.LoggerMW())

	return &HttpRouter{router: router}
}

func (r *HttpRouter) Register(handler Handler) {
	r.router.GET("/banners", handler.GetBanners)
	r.router.POST("/banners/create", handler.CreateBanner)
	r.router.PUT("/banners/update", handler.UpdateBanner)
	r.router.DELETE("/banners/delete", handler.DeleteBanner)
}

func (r *HttpRouter) Run(address string) error {
	return r.router.Run(address)
}
