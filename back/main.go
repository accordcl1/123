package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// Структура ответа в формате JSON
type ApiResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	RandomMetric int    `json:"random_metric"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Указываем заголовок, что отдаем JSON
	w.Header().Set("Content-Type", "application/json")

	response := ApiResponse{
		Status:       "success",
		Message:      "Привет из Go-бэкенда!",
		RandomMetric: rand.Intn(491) + 10, // Случайное число от 10 до 500
	}

	// Кодируем структуру в JSON и отправляем клиенту
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", apiHandler)
	
	// Запускаем сервер на порту 5000
	http.ListenAndServe(":5000", nil)
}
