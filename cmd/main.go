package main

import (
	"fmt"
	"log"
	"mygo/internal" // Импорт пакета с InitDB
	"net/http"
)

func main() {
	// Инициализация подключения к БД
	dataSourceName := "host=localhost port=5432 user=postgres password=2333 dbname=postgres sslmode=disable"
	internal.InitDB(dataSourceName)

	// Регистрация обработчиков
	http.HandleFunc("/banners", internal.GetBannersHandler)           // Получение всех баннеров
	http.HandleFunc("/banners/create", internal.CreateBannerHandler)  // Создание нового баннера
	http.HandleFunc("/banners/update/", internal.UpdateBannerHandler) // Обновление баннера
	http.HandleFunc("/banners/delete/", internal.DeleteBannerHandler) // Удаление баннера

	// Запуск сервера
	log.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err) // Логируем ошибку и завершаем программу
	}

	fmt.Println("Сервер запущен...")
}
