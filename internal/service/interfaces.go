package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
)

type Service struct {
	TestService         TestService
	IncidentService     IncidentService
	IncidentTypeService IncidentTypeService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TestService:         NewTestService(repos.TestRepository),
		IncidentService:     NewIncidentService(repos.IncidentRepository, repos.IncidentTypeRepository),
		IncidentTypeService: NewIncidentTypeService(repos.IncidentTypeRepository),
	}
}
