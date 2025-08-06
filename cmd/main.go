package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// Создание нового логгера
	logger := log.New(os.Stdout, "Обработчик кода Морзе: ", log.Ldate|log.Ltime)

	// Создание нового сервера
	server := server.NewServer(logger)

	// Запуск сервера
	logger.Printf("Старт сервера на %s ", server.HttpServer.Addr)

	err := http.ListenAndServe(server.HttpServer.Addr, server.HttpServer.Handler)

	// Обработка ошибок
	if err != nil {
		logger.Fatalf("Критическая ошибка: %v", err)
	}
}
