package repository

import (
	"github.com/txzy2/go-logger-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IncidentTypeRepository interface {
	FindByCode(code string) (bool, error)
}

type incidentTypeRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewIncidentTypeRepository(logger *zap.Logger, db *gorm.DB) IncidentTypeRepository {
	return &incidentTypeRepository{
		logger: logger,
		db:     db,
	}
}

func (r *incidentTypeRepository) FindByCode(code string) (bool, error) {
	err := r.db.Model(&models.IncidentType{}).Where("code = ?", code).First(&models.IncidentType{}).Error
	if err != nil {
		r.logger.Error("Error finding incident type", zap.Error(err))
		return false, err
	}
	return true, nil
}
