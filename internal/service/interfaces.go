package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	TestService         TestService
	IncidentService     IncidentService
	IncidentTypeService IncidentTypeService
	TemplateService     TemplateService
	logger              *zap.Logger
}

func NewService(repos *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		TestService:         NewTestService(logger, repos.TestRepository),
		IncidentService:     NewIncidentService(logger, repos.IncidentRepository, repos.IncidentTypeRepository, repos.TemplateRepository),
		IncidentTypeService: NewIncidentTypeService(logger, repos.IncidentTypeRepository),
		TemplateService:     NewTemplateService(logger, repos.TemplateRepository),
		logger:              logger,
	}
}
