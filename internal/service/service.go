package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndConvert определяет тип входной строки (Морзе или текст) и конвертирует её.
func DetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("input is empty")
	}
	//
	// Если строка содержит только точки, тире и пробелы — вероятно, это Морзе.
	if IsMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	// Иначе — обычный текст.
	return morse.ToMorse(trimmed), nil
}

// isMorse проверяет, является ли строка кодом Морзе.
func IsMorse(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return !(r == '.' || r == '-' || r == ' ' || r == '\n')
	})
}
