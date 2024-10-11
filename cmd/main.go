package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mygo/internal/db"
	"mygo/internal/handler"
	"mygo/internal/repository"
	"mygo/internal/router"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=2333 dbname=postgres sslmode=disable"

	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	repo := repository.NewRepository(database)
	handler := handler.NewHandler(repo)

	r := gin.Default()

	router.RegisterRoutes(r, handler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
