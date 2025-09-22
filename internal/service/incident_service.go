package service

import (
	"errors"
	"fmt"

	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/parsers"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type IncidentService interface {
	WriteOrSaveLogs(data types.IncidentData) string
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

func (s *incidentService) WriteOrSaveLogs(data types.IncidentData) string {
	parseData, err := s.parseIncidentMessage(data)
	if err == nil {
		res, err := s.incidentTypeRepo.FindByCode(parseData.Code)
		if err == nil {
			s.logger.Warn("Incident type retrieved",
				zap.Any("incident_type", res),
				zap.String("method", "GetIncidentType"),
			)
			return "SUCCESS"
		}
	}

	return fmt.Sprintf("Error finding incident type: %v", err)
}

func (s *incidentService) parseIncidentMessage(data types.IncidentData) (parsers.ParserMessageResponse, error) {
	parser, err := parsers.NewParser(data.Service, data)
	s.logger.Info("Parser created")
	if err != nil {
		return parsers.ParserMessageResponse{}, errors.New("Error creating parser for service")
	}

	parseData, err := parser.ParseMessage(data.Message)
	if err != nil {
		return parsers.ParserMessageResponse{}, errors.New("Error parsing data for service")
	}
	s.logger.Info("Data parsed", zap.Any("data", parseData))

	return parseData, nil
}
