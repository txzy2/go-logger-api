package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
	"go.uber.org/zap"
)

type IncidentTypeService interface{}

type incidentTypeService struct {
	logger           *zap.Logger
	incidentTypeRepo repository.IncidentTypeRepository
}

func NewIncidentTypeService(logger *zap.Logger, incidentTypeRepo repository.IncidentTypeRepository) IncidentTypeService {
	return &incidentTypeService{
		incidentTypeRepo: incidentTypeRepo,
		logger:           logger,
	}
}
