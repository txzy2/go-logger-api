package repository

import (
	"log"

	"github.com/txzy2/go-logger-api/internal/models"
	"gorm.io/gorm"
)

type IncidentTypeRepository interface {
	FindByCode(code string) (bool, error)
}

type incidentTypeRepository struct {
	db *gorm.DB
}

func NewIncidentTypeRepository(db *gorm.DB) IncidentTypeRepository {
	return &incidentTypeRepository{
		db: db,
	}
}

func (r *incidentTypeRepository) FindByCode(code string) (bool, error) {
	err := r.db.Model(&models.IncidentType{}).Where("code = ?", code).First(&models.IncidentType{}).Error
	if err != nil {
		log.Printf("Error finding incident type: %v", err)
		return false, err
	}
	return true, nil
}
