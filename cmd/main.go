package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// Создание нового логгера
	logger := log.New(os.Stdout, "Обработчик кода Морзе: ", log.Ldate|log.Ltime)

	// Создание нового сервера
	srv := server.NewServer(logger)

	// Запуск сервера
	logger.Printf("Старт сервера на %s ", srv.HttpServer.Addr)

	err := srv.HttpServer.ListenAndServe()

	// Обработка ошибок
	if err != nil {
		logger.Fatalf("Критическая ошибка: %v", err)
	}
}
