package service

import (
	"github.com/txzy2/go-logger-api/internal/models"
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

	s.validateSenderMethod(incidentType)
}

func (s *incidentService) validateSenderMethod(incidentType *models.IncidentType) {
	if incidentType.Alias == "" {
		s.logger.Warn("Incident type is not supported", zap.Any("data", incidentType), zap.String("method", "FindByCode"))
		return
	}

	//TODO: Сделать получение sendTemplate по id из таблицы по send_template_id

}
