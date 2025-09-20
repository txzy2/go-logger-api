package models

import (
	"time"

	"gorm.io/gorm"
)

// IncidentType представляет таблицу incident_type
type IncidentType struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeName       string         `gorm:"type:varchar(50);not null" json:"type_name"`
	SendTemplateID uint           `gorm:"not null" json:"send_template_id"`
	Code           string         `gorm:"type:varchar(50);not null" json:"code"`
	Lifecycle      *string        `gorm:"type:varchar(255)" json:"lifecycle,omitempty"`
	Alias          string         `gorm:"type:varchar(255);default:'manager';not null;check:alias IN ('manager','client')" json:"alias"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Связи
	SendTemplate SendTemplate `gorm:"foreignKey:SendTemplateID;references:ID" json:"send_template,omitempty"`
	Incidents    []Incident   `gorm:"foreignKey:IncidentTypeID;references:ID" json:"incidents,omitempty"`
}

// TableName возвращает имя таблицы для модели IncidentType
func (IncidentType) TableName() string {
	return "incident_type"
}
