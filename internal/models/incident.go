package models

import (
	"time"

	"gorm.io/datatypes"
)

// Incident представляет таблицу incident
// @Description Модель инцидента в системе логирования
type Incident struct {
	ID               uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Service          string         `gorm:"type:varchar(255);not null" json:"service"`
	Level            string         `gorm:"type:varchar(255);not null" json:"level"`
	Message          string         `gorm:"type:varchar(255);not null" json:"message"`
	Action           string         `gorm:"type:varchar(255);not null" json:"action"`
	AdditionalFields datatypes.JSON `gorm:"type:json;not null" json:"additionalFields"`
	Function         string         `gorm:"type:varchar(255);not null" json:"function"`
	Class            string         `gorm:"type:varchar(255);not null" json:"class"`
	File             string         `gorm:"type:varchar(255);not null" json:"file"`
	Date             time.Time      `gorm:"type:date;not null" json:"date"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

// TableName возвращает имя таблицы для модели Incident
func (Incident) TableName() string {
	return "incident"
}
