package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
	"go.uber.org/zap"
)

type TemplateService interface {
}

type templateService struct {
	templateRepo repository.TemplateRepository
	logger       *zap.Logger
}

func NewTemplateService(logger *zap.Logger, templateRepository repository.TemplateRepository) TemplateService {
	return &templateService{
		templateRepo: templateRepository,
		logger:       logger,
	}
}
