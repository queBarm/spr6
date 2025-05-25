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
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг multipart формы (10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}

	// Получаем файл по имени "file"
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение содержимого
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	input := string(content)

	// Обработка содержимого
	converted, err := service.DetectAndConvert(input)
	if err != nil {
		http.Error(w, "Failed to convert content", http.StatusBadRequest)
		return
	}
	// Создание временного файла
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("converted_%s%s", time.Now().UTC().Format("20060102T150405"), ext)
	outputPath := filepath.Join(os.TempDir(), filename)

	outFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Запись преобразованного текста в файл
	_, err = outFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	// Корректный порядок: сначала заголовки
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
