package service

import (
	"errors"

	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/parsers"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type IncidentService interface {
	WriteOrSaveLogs(data types.IncidentData)
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

func (s *incidentService) WriteOrSaveLogs(data types.IncidentData) {
	parseData, err := s.parseIncidentMessage(data)
	if err != nil {
		s.logger.Warn("Incident type retrieved", zap.Error(err), zap.Any("data", data), zap.String("method", "parseIncidentMessage"))
		return
	}

	res, err := s.incidentTypeRepo.FindByCode(parseData.Code)
	if err != nil {
		s.logger.Warn("Incident type retrieved", zap.Error(err), zap.Any("Code", parseData.Code), zap.String("method", "FindByCode"))
		return
	}

	s.incidentRepo.CreateIncident(data, res.ID)
}

func (s *incidentService) parseIncidentMessage(data types.IncidentData) (parsers.ParserMessageResponse, error) {
	parser, err := parsers.NewParser(data.Service, data)
	if err != nil {
		return parsers.ParserMessageResponse{}, errors.New("Error creating parser for service")
	}

	parseData, err := parser.ParseMessage(data.Message)
	if err != nil {
		return parsers.ParserMessageResponse{}, errors.New("Error parsing data for service")
	}

	return parseData, nil
}
