package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/parsers"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type IncidentService interface {
	WriteOrSaveLogs(data types.IncidentData) string
}

type incidentService struct {
	incidentRepo     repository.IncidentRepository
	incidentTypeRepo repository.IncidentTypeRepository
}

func NewIncidentService(
	incidentRepo repository.IncidentRepository,
	incidentTypeRepo repository.IncidentTypeRepository,
) IncidentService {
	return &incidentService{
		incidentRepo:     incidentRepo,
		incidentTypeRepo: incidentTypeRepo,
	}
}

func (s *incidentService) WriteOrSaveLogs(data types.IncidentData) string {
	parsedData, err := s.parseIncidentMessage(data)
	if err == nil {
		res, err := s.incidentTypeRepo.FindByCode(parsedData["code"])
		if err == nil {
			log.Printf("Incident type: %v", res)
			return "SUCCESS"
		}
	}

	return fmt.Sprintf("Error finding incident type: %v", err)
}

func (s *incidentService) parseIncidentMessage(data types.IncidentData) (map[string]string, error) {
	parser, err := parsers.NewParser(data.Service, data)
	log.Printf("parser: %v", parser)
	if err != nil {
		return nil, errors.New("Error creating parser for service")
	}

	parsed, err := parser.ParseMessage(data.Message)
	if err != nil {
		return nil, errors.New("Error parsing data for service")
	}
	log.Printf("Parsed data: %v", parsed)

	return parsed, nil
}
