package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Обработка данных из файла
func ConvertPayload(s string) (string, error) {

	// Проверка на то, что данные содержат код Морзе
	isMorse, err := IsMorse(s)

	if err != nil {
		return "", err
	}

	if isMorse {
		// Конвертация кода Морзе в текст
		return morse.ToText(s), nil
	}

	// Конвертация текста в код Морзе
	return morse.ToMorse(s), nil
}

// Проверка того, что строка является кодом Морзе
func IsMorse(s string) (bool, error) {

	// Ошибки выполнения
	var (
		errEmptyString = errors.New("String is empty")
	)

	// Если строка пустая, то возвращаем ошибку
	if s == "" {
		return false, errEmptyString
	}

	// Для проверки строки переводим все символы в нижний регистр
	s = strings.ToLower(s)

	// Строка не является кодом Морзе, если содержит хотя бы один символ, отличающийся от . и -
	if strings.IndexAny(s, "абвгдеёжзийклмнопрстуфхцчшщъыьэюя0123456789,:?\\/()\"") == -1 {
		return true, nil
	}

	return false, nil
}
