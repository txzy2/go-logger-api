package repository

import (
	"errors"
	"fmt"

	"github.com/txzy2/go-logger-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TemplateRepository interface {
	FindByID(id uint) (*models.SendTemplate, error)
}

type templateRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewTemplateRepository(logger *zap.Logger, db *gorm.DB) TemplateRepository {
	return &templateRepository{
		logger: logger,
		db:     db,
	}
}

func (tm *templateRepository) FindByID(id uint) (*models.SendTemplate, error) {
	var template models.SendTemplate

	if err := tm.db.First(&template, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tm.logger.Warn("Template not found", zap.Uint("id", id))
			return nil, fmt.Errorf("template with ID '%d' not found", id)
		}
		return nil, err
	}
	return &template, nil
}
