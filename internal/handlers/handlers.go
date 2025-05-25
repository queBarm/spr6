package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
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
	isMorse := service.IsMorse(input)

	var result string
	if isMorse {
		result = morse.ToText(input)
	} else {
		result = morse.ToMorse(input)
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
