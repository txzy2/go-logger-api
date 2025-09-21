package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
)

type IncidentService interface{}

type incidentService struct {
	incidentRepo repository.IncidentRepository
}

func NewIncidentService(incidentRepo repository.IncidentRepository) IncidentService {
	return &incidentService{
		incidentRepo: incidentRepo,
	}
}
