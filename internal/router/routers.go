package router

import (
	"github.com/gin-gonic/gin"

	"mygo/internal/handler"
)

func RegisterRoutes(r *gin.Engine, handler *handler.Handler) {
	r.GET("/banners", handler.GetBanners)
	r.POST("/banners/create", handler.CreateBanner)
	r.PUT("/banners/update", handler.UpdateBanner)
	r.DELETE("/banners/delete", handler.DeleteBanner)
}
