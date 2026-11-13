package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// Структура ответа в формате JSON с добавленным полем Version
type ApiResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	RandomMetric int    `json:"random_metric"`
	Version      string `json:"version"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Указываем заголовок, что отдаем JSON
	w.Header().Set("Content-Type", "application/json")

	response := ApiResponse{
		Status:       "success",
		Message:      "Привет из Go-бэкенда (v2)!",
		RandomMetric: rand.Intn(491) + 10, // Случайное число от 10 до 500
		Version:      "1.0.1",
	}

	// Кодируем структуру в JSON и отправляем клиенту
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", apiHandler)
	
	// Запускаем сервер на порту 5000 внутри контейнера
	http.ListenAndServe(":5000", nil)
}
