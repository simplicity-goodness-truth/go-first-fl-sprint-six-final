package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Путь к корневому index.html
const indexHTML = "../index.html"

// Обработка корневого эндпоинта /: возвращаем HTML из файла
func HandleRoot(res http.ResponseWriter, req *http.Request) {

	// Установка типа содержимого text/html и успешного статуса
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	// Разрешен только метод GET
	if req.Method != http.MethodGet {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Чтение содержимого index.html
	rootContents, err := readFile(indexHTML)

	if err != nil {
		http.Error(res, "Failed to open html file:"+err.Error(), http.StatusBadRequest)
		return
	}

	// Вывод содержимого index.html в ответ сервера
	res.Write(rootContents)
}

// Обработка эндпоинта /upload
func HandleUpload(res http.ResponseWriter, req *http.Request) {

	// Установка типа содержимого text/html; charset=utf-8 и успешного статуса
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	// Разрешен только метод POST
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получение загруженного файла
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Failed to get file:"+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Чтение содержимого загруженного файла
	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Failed to read file content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Конвертация содержимого загруженного файла в строку
	payload := string(fileContent)

	// Конвертация содержимого в текст или код Морзе
	convertedPayload, err := service.ConvertPayload(payload)

	if err != nil {

		// Вывод информации об ошибке при неуспешной конвертации

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Failed to convert file content:" + err.Error()))

		return
	}

	// Подготовка названия файла из временной отметки UTC и расширения log

	fileName := fmt.Sprintf("%s%s", time.Now().UTC().Format("2006-01-02_15-04-05") , filepath.Ext("*.txt"))

	// Запись результата конвертации в файл
	err = writeStringToFile(fileName, convertedPayload)

	if err != nil {
		http.Error(res, "Failed to record file content: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Вывод результата конвертации в ответ сервера
	res.Write([]byte(convertedPayload))
}

// Чтения файла по его пути
func readFile(filePath string) ([]byte, error) {

	// Открытие файла на чтение
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)

	defer file.Close()

	if err != nil {
		return nil, err
	}

	// Чтение содержимого файла
	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	// Возвращаем содержимое файла
	return data, nil
}

// Запись в файл строки по его пути
func writeStringToFile(filePath string, content string) error {

	// Открытие файла на запись

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		return err
	}

	return nil
}
