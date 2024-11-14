package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"mygo/internal/auth"
)

// /////////////////////////////////////////////////////////////////////////
func JWTMiddleware(c *gin.Context) {
	// Извлекаем токен из заголовка Authorization
	tokenStr := c.GetHeader("Authorization")
	// Убираем префикс "Bearer "
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// Проверяем токен
	claims, err := auth.ValidateJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		c.Abort() // Прерываем обработку, если токен невалидный
		return
	}

	// Добавляем данные пользователя в контекст запроса
	c.Set("username", claims.Username)

	// Переходим к следующему обработчику
	c.Next()
}

////////////////////////////////////////////////////////////////////////////

func LoggerMiddleware() gin.HandlerFunc {
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
