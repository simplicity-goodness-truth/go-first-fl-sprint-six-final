package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// Создание нового логгера
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime)

	// Создание нового сервера
	server := server.NewServer(errorLog)

	// Запуск сервера
	err := http.ListenAndServe(server.HttpServer.Addr, server.HttpServer.Handler)

	// Обработка ошибок
	if err != nil {
		errorLog.Fatal(err)
	}
}
