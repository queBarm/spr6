package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler обрабатывает GET / и возвращает HTML-страницу.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

// UploadHandler обрабатывает POST /upload, читает файл, конвертирует, сохраняет, возвращает результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Максимальный размер формы, например, 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	// Получаем файл
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file contents", http.StatusInternalServerError)
		return
	}

	// Конвертируем через service
	result, err := service.DetectAndConvert(string(data))
	if err != nil {
		http.Error(w, "Failed to convert input", http.StatusInternalServerError)
		return
	}

	// Формируем имя для сохранения
	ext := filepath.Ext(handler.Filename)
	timestamp := time.Now().UTC().Format("20060102_150405")
	outFilename := fmt.Sprintf("converted_%s%s", timestamp, ext)

	// Создаем файл
	outFile, err := os.Create(outFilename)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Записываем результат
	if _, err := outFile.WriteString(result); err != nil {
		http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат клиенту
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(result))
}
