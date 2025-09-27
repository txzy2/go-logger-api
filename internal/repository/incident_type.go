package repository

import (
	"errors"
	"fmt"

	"github.com/txzy2/go-logger-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IncidentTypeRepository interface {
	FindByCode(code string) (*models.IncidentType, error)
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

func (r *incidentTypeRepository) FindByCode(code string) (*models.IncidentType, error) {
	var incidentType models.IncidentType
	err := r.db.Preload("SendTemplate").Where("code = ?", code).First(&incidentType).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warn("Incident type not found", zap.String("code", code))
			return nil, fmt.Errorf("incident type with code '%s' not found", code)
		}

		r.logger.Error("Error finding incident type", zap.Error(err))
		return nil, err
	}

	return &incidentType, nil
}
