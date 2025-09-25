package service

import (
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/parsers"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type IncidentService interface {
	ProcessIncident(data types.IncidentData)
}

type incidentService struct {
	logger           *zap.Logger
	incidentRepo     repository.IncidentRepository
	incidentTypeRepo repository.IncidentTypeRepository
}

func NewIncidentService(
	logger *zap.Logger,
	incidentRepo repository.IncidentRepository,
	incidentTypeRepo repository.IncidentTypeRepository,
) IncidentService {
	return &incidentService{
		logger:           logger,
		incidentRepo:     incidentRepo,
		incidentTypeRepo: incidentTypeRepo,
	}
}

func (s *incidentService) ProcessIncident(data types.IncidentData) {
	s.logger.Info("Incident data", zap.Any("data", data), zap.String("method", "CreateIncident"))
	// Пытаемся записать инцидент
	res := s.incidentRepo.CreateIncident(data)
	if !res {
		return
	}

	// Ищем поле code в AdditionalField чтобы потом попробовать отпаравить
	code, ok := parsers.FindKeyInAdditionalFields(data.AdditionalFields, "code")
	if !ok {
		s.logger.Warn("Incident AdditionalField find key error", zap.Any("data", data), zap.String("method", "hasKey"))
		return
	}

	// Проверяем отслеживаем ли такой тип ошибки
	incidentType, err := s.incidentTypeRepo.FindByCode(code)
	if err != nil {
		s.logger.Warn("Incident Code is not found", zap.Error(err), zap.Any("Code", code), zap.String("method", "FindByCode"))
		return
	}
	s.logger.Info("Incident type retrieved", zap.Any("data", incidentType), zap.String("method", "FindByCode"))

	// TODO: Сделать отправку (фильтровать по email или push)

}
