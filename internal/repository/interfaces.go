package repository

import (
	"github.com/txzy2/go-logger-api/pkg/database"
	"go.uber.org/zap"
)

type Repository struct {
	logger                 *zap.Logger
	TestRepository         TestRepository
	IncidentRepository     IncidentRepository
	IncidentTypeRepository IncidentTypeRepository
	TemplateRepository     TemplateRepository
}

func NewRepository(logger *zap.Logger, db *database.Database) *Repository {
	return &Repository{
		logger:                 logger,
		TestRepository:         NewTestRepository(logger, db.GORM),
		IncidentRepository:     NewIncidentRepository(logger, db.GORM),
		IncidentTypeRepository: NewIncidentTypeRepository(logger, db.GORM),
		TemplateRepository:     NewTemplateRepository(logger, db.GORM),
	}
}
