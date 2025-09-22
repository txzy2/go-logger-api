package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	TestService         TestService
	IncidentService     IncidentService
	IncidentTypeService IncidentTypeService
	logger              *zap.Logger
}

func NewService(repos *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		TestService:         NewTestService(logger, repos.TestRepository),
		IncidentService:     NewIncidentService(logger, repos.IncidentRepository, repos.IncidentTypeRepository),
		IncidentTypeService: NewIncidentTypeService(logger, repos.IncidentTypeRepository),
		logger:              logger,
	}
}
