package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMW() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("до")
		ctx.Next() //говно
		log.Println("после")

		log.Printf("Тип и путь: %s %s      | Статус: %d     |\n",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Writer.Status(),
		)
	}
}
