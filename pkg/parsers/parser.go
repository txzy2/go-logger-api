package parsers

import (
	"time"

	"github.com/txzy2/go-logger-api/pkg/types"
)

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func FindKeyInAdditionalFields(fields []types.AdditionalField, key string) (string, bool) {
	for _, field := range fields {
		if field.Key == key {
			return field.Value, true
		}
	}
	return "", false
}
