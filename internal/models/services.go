package models

import (
	"time"
)

type ActiveEnum string

const (
	ActiveStatus   ActiveEnum = "Y"
	InactiveStatus ActiveEnum = "N"
)

// Services представляет таблицу services
// @Description Модель сервисов в системе
type Services struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Active    ActiveEnum `gorm:"type:varchar(255);not null" json:"active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (Services) TableName() string {
	return "services"
}
