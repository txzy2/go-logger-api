package repository

import (
	"errors"
	"fmt"

	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/pkg/types"
	"gorm.io/gorm"
)

type IncidentRepository interface {
	FindByName(serviceName types.Service) (bool, error)
}

type incidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) IncidentRepository {
	return &incidentRepository{
		db: db,
	}
}

func (r *incidentRepository) FindByName(serviceName types.Service) (bool, error) {
	err := r.db.Model(&models.Services{}).Where("name = ? AND active = ?", serviceName, models.ActiveStatus).First(&models.Services{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("Сервис не найден или не доступен")
		}
	}
	return err == nil, err
}
