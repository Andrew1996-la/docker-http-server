package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

// Start запускает HTTP-сервер
func Start() error {

	// Регистрируем handler для пути /ping
	// Когда приходит запрос на /ping — выполняется эта функция
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		// Лог в консоль (на стороне сервера)
		fmt.Println("processed request /ping")

		// Отправляем ответ клиенту
		w.Write([]byte("hello from Docker\n"))
	})

	// Запускаем HTTP-сервер на порту 5050
	err := http.ListenAndServe(":5050", nil)

	// Если сервер был остановлен корректно (например через shutdown)
	// — это НЕ ошибка, просто выходим
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	// Любая другая ошибка — настоящая проблема
	return err
}
