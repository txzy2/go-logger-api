package models

import (
	"time"

	"gorm.io/gorm"
)

// SendTemplate представляет таблицу send_template
// @Description Модель шаблона для отправки уведомлений
type SendTemplate struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	To        string         `gorm:"type:varchar(255);not null" json:"to"`
	Subject   string         `gorm:"type:varchar(255);not null" json:"subject"`
	Template  string         `gorm:"type:text;not null" json:"template"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

// TableName возвращает имя таблицы для модели SendTemplate
func (SendTemplate) TableName() string {
	return "send_template"
}
