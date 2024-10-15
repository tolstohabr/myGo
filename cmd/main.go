package main

import (
	"flag"
	"log"

	"mygo/internal/config"
	"mygo/internal/db"
	"mygo/internal/handler"
	"mygo/internal/repository"
	"mygo/internal/router"
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
	handler := handler.NewHandler(repo)

	router := router.NewHttpRouter()
	router.Register(handler)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("sdfsdsd", err)
	}
	//TODO: "localhost:8080" из кофниг

	//TODO: сервисы. хендлер вызыввает серавис и наоборот. Сделать у них интерфейсы(у хендлера интерфейс сервиса, а у сервиса интерфейс репозитория)

	//TODO: аутентификация в Go посмотреть (как с мидлвере)
}
