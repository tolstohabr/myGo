package router

import (
	"mygo/internal/handler"

	"github.com/gin-gonic/gin"
)

// Функция регистрации маршрутов в Gin
func RegisterRoutes(r *gin.Engine, handler *handler.Handler) {
	r.GET("/banners", handler.GetBanners)             // Регистрируем GET
	r.POST("/banners/create", handler.CreateBanner)   // Регистрируем POST
	r.PUT("/banners/update", handler.UpdateBanner)    // Регистрируем PUT
	r.DELETE("/banners/delete", handler.DeleteBanner) // Регистрируем DELETE
}
