package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура для сервера
type Server struct {
	Logger     *log.Logger
	HttpServer *http.Server
}

// Функция для создания http-роутера
func NewServer(logger *log.Logger) *Server {

	// Регистрация хэндлеров
	router := http.NewServeMux()
	router.HandleFunc("/upload", handlers.HandleUpload)
	router.HandleFunc("/", handlers.HandleRoot)

	// Подготовка экземляра структуры сервера

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HttpServer: server,
	}
}
