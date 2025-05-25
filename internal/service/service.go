package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndConvert определяет тип входной строки (Морзе или текст) и конвертирует её.
func DetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Если строка содержит только точки, тире и пробелы — вероятно, это Морзе.
	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	// Иначе — обычный текст.
	return morse.ToMorse(trimmed), nil
}

// isMorse проверяет, является ли строка кодом Морзе.
func isMorse(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '\t', '\n', '\r':
			// допустимые символы в морзе
		default:
			if !unicode.IsSpace(r) {
				return false
			}
		}
	}
	return true

}
