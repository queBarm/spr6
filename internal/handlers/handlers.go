package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает GET / и возвращает HTML-страницу.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает POST /upload и возвращает конвертированный результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничиваем только POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Чтение файла из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Считываем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}

	input := string(content)

	// Определяем тип содержимого (морзе или текст)
	result, err := service.DetectAndConvert(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(result)); err != nil {
		log.Printf("Write respone error: %v", err)
	}
}
