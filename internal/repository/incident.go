package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type IncidentRepository interface {
	FindByName(serviceName types.Service) (bool, error)
	CreateIncident(data types.IncidentData, incidentTypeId uint)
}

type incidentRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewIncidentRepository(logger *zap.Logger, db *gorm.DB) IncidentRepository {
	return &incidentRepository{
		logger: logger,
		db:     db,
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

func (r *incidentRepository) CreateIncident(data types.IncidentData, incidentTypeId uint) {
	user := &models.Incident{
		Service:          string(data.Service),
		Level:            data.Level,
		Message:          data.Message,
		IncidentTypeID:   incidentTypeId,
		Action:           data.Action,
		AdditionalFields: r.formatAdditionalFields(&data.AdditionalFields),
		Function:         data.Function,
		Class:            data.Class,
		File:             data.File,
		Date:             data.Date,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := r.db.Create(&user).Error
	if err != nil {
		r.logger.Error("Error creating incident", zap.Error(err))
	}

	r.logger.Info("Incident created", zap.Any("incident", user))
}

func (r *incidentRepository) formatAdditionalFields(additionalFields *[]types.AdditionalField) string {
	var formattedFields []string
	for _, field := range *additionalFields {
		formattedFields = append(formattedFields, fmt.Sprintf("%s: %s", field.Key, field.Value))
	}

	return strings.Join(formattedFields, ", ")
}
