package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mygo/internal/config"
	"mygo/internal/db"
	"mygo/internal/handler"
	"mygo/internal/repository"
	"mygo/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	database, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	repo := repository.NewRepository(database)
	handler := handler.NewHandler(repo)

	r := gin.Default()

	r.Use(LoggerMW())

	router.RegisterRoutes(r, handler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func LoggerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("до")
		c.Next() //я не понимаю как эта херня работет
		log.Println("после")

		log.Printf("Тип и путь: %s %s      | Статус: %d     |\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
		)
	}
}
