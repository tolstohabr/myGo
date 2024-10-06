package main

import (
	"fmt"
	"log"

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
	handler := handler.NewHander(repo)

	r := router.NewRouter()
	r.Register(handler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}

	fmt.Println("подвал")
}
