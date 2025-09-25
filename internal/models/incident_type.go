package models

import (
	"time"
)

// IncidentType представляет таблицу incident_type
// @Description Модель типа инцидента
type IncidentType struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeName       string    `gorm:"type:varchar(50);not null" json:"type_name"`
	SendTemplateID uint      `gorm:"not null" json:"send_template_id"`
	Code           string    `gorm:"type:varchar(50);not null" json:"code"`
	Lifecycle      *string   `gorm:"type:varchar(255)" json:"lifecycle,omitempty"`
	Alias          string    `gorm:"type:varchar(255);default:'NULL';check:alias IN ('email','push')" json:"alias"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Связи
	SendTemplate SendTemplate `gorm:"foreignKey:SendTemplateID;references:ID" json:"send_template,omitempty"`
}

// TableName возвращает имя таблицы для модели IncidentType
func (IncidentType) TableName() string {
	return "incident_type"
}
