package internal

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Подключение драйвера PostgreSQL
)

// DB является глобальным объектом базы данных
var DB *sql.DB

// InitDB инициализирует подключение к базе данных
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
