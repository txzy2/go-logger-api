package types

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// AdditionalField представляет дополнительное поле инцидента
// @Description Структура для хранения дополнительных полей инцидента в формате ключ-значение
type AdditionalField struct {
	Key   string `json:"key" example:"user_id" description:"Ключ дополнительного поля"`
	Value string `json:"value" example:"12345" description:"Значение дополнительного поля"`
}

var validate = validator.New()

// IncidentData представляет данные инцидента для создания
// @Description Структура данных для создания нового инцидента в системе
// TODO: Убрать привязку к сервисам
type IncidentData struct {
	Level            string            `json:"level" validate:"required" example:"ERROR" description:"Уровень логирования (error, warning, info)"`
	Message          string            `json:"message" validate:"required" example:"Database connection failed" description:"Сообщение об ошибке или событии"`
	Domain           string            `json:"domain" validate:"required" example:"database" description:"Домен системы, где произошло событие"`
	Action           string            `json:"action" validate:"required" example:"connect" description:"Действие, которое выполнялось при возникновении события"`
	Function         string            `json:"function" validate:"required" example:"ConnectDB" description:"Название функции, где произошло событие"`
	Service          string            `json:"service" validate:"required" example:"user-service" description:"Название сервиса"`
	File             string            `json:"file" validate:"required" example:"main.go" description:"Файл, где произошло событие"`
	Class            string            `json:"class" validate:"required" example:"User" description:"Класс, где произошло событие"`
	AdditionalFields []AdditionalField `json:"additionalFields" validate:"required" description:"Дополнительные поля с метаданными"`
	Date             time.Time         `json:"date" validate:"required" example:"2024-01-15T10:30:00Z" description:"Дата и время события в ISO 8601 формате"`
}

func ValidateIncidentData(data IncidentData) error {
	err := validate.Struct(data)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}
