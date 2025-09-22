package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
)

type IncidentTypeService interface{}

type incidentTypeService struct {
	incidentTypeRepo repository.IncidentTypeRepository
}

func NewIncidentTypeService(incidentTypeRepo repository.IncidentTypeRepository) IncidentTypeService {
	return &incidentTypeService{
		incidentTypeRepo: incidentTypeRepo,
	}
}
