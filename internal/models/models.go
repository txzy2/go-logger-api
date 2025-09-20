package models

import (
	"gorm.io/gorm"
)

// AllModels возвращает список всех моделей для миграции
func AllModels() []interface{} {
	return []interface{}{
		&SendTemplate{},
		&IncidentType{},
		&Incident{},
		&Services{},
	}
}

// AutoMigrate выполняет автоматическую миграцию всех моделей
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(AllModels()...)
}

// DropTables удаляет все таблицы (используется для тестов)
func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(AllModels()...)
}
