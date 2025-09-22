package models

import (
	"time"
)

// Incident представляет таблицу incident
// @Description Модель инцидента в системе логирования
type Incident struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Service          string    `gorm:"type:varchar(255);not null" json:"service"`
	Level            string    `gorm:"type:varchar(255);not null" json:"level"`
	Message          string    `gorm:"type:varchar(255);not null" json:"message"`
	IncidentTypeID   uint      `gorm:"not null" json:"incident_type_id"`
	Action           string    `gorm:"type:varchar(255);not null" json:"action"`
	AdditionalFields string    `gorm:"type:text;not null" json:"additionalFields"`
	Function         string    `gorm:"type:varchar(255);not null" json:"function"`
	Class            string    `gorm:"type:varchar(255);not null" json:"class"`
	File             string    `gorm:"type:varchar(255);not null" json:"file"`
	Date             time.Time `gorm:"type:date;not null" json:"date"`
	Count            int       `gorm:"not null" json:"count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// Связи
	IncidentType IncidentType `gorm:"foreignKey:IncidentTypeID;references:ID" json:"incident_type,omitempty"`
}

// TableName возвращает имя таблицы для модели Incident
func (Incident) TableName() string {
	return "incident"
}
