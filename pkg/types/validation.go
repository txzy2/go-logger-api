package types

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError форматирует ошибки валидации в понятный вид
// Можно переиспользовать для любых структур с валидацией
func FormatValidationError(err error) error {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		errors = append(errors, fmt.Sprintf("%s is required", field))
	}
	return fmt.Errorf(strings.Join(errors, ", "))
}
