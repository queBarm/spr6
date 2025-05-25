package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// Создаем сервер через пакет server
	srv := server.NewServer(logger)

	// Запускаем сервер
	logger.Println("Запуск сервера на http://localhost:8080")
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
