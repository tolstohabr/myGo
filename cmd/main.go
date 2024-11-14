package main

import (
	"flag"
	"log"

	"mygo/internal/config"
	"mygo/internal/db"
	"mygo/internal/handler"
	"mygo/internal/repository"
	"mygo/internal/router"
	"mygo/internal/service"
)

var configPath = flag.String("config", "./config/config.yaml", "config path")

func main() {
	flag.Parse()

	cfg := config.MustLoad(*configPath)

	database, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	repo := repository.NewRepository(database)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := router.NewHttpRouter()
	router.Register(handler)
	if err := router.Run(cfg.HOST); err != nil {
		log.Fatal("fdfdf", err)
	}
}
