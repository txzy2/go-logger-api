package service

import (
	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/helpers/senders"
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
	templateRepo     repository.TemplateRepository
}

func NewIncidentService(
	logger *zap.Logger,
	incidentRepo repository.IncidentRepository,
	incidentTypeRepo repository.IncidentTypeRepository,
	templateRepo repository.TemplateRepository,
) IncidentService {
	return &incidentService{
		logger:           logger,
		incidentRepo:     incidentRepo,
		incidentTypeRepo: incidentTypeRepo,
		templateRepo:     templateRepo,
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

	s.validateAndSendToRecipient(incidentType, &data)
}

func (s *incidentService) validateAndSendToRecipient(incidentType *models.IncidentType, incidentData *types.IncidentData) {
	if incidentType.Alias == "" {
		s.logger.Warn("Incident type is not supported", zap.Any("data", incidentType), zap.String("method", "FindByCode"))
		return
	}

	// Создвем нового отправителя
	sender, err := senders.NewSenderManager(incidentType.Alias, incidentType, incidentData, s.logger)
	if err != nil {
		s.logger.Error("Error creating sender", zap.Error(err))
		return
	}

	res, err := sender.PrepareIncidentData()
	if err != nil {
		s.logger.Error("Error preparing incident data for send", zap.Error(err))
		return
	}

	req := sender.Send(res.Emails)
	s.logger.Info("Incident sent", zap.Bool("data", req), zap.String("method", "validateAndSendToRecipient"))

	//TODO: Сделать запись в новую таблицу SendStatus
}
