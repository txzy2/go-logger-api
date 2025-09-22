package repository

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TestRepository interface {
	Ping(ctx context.Context) error
}

type testRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

// NewTestRepository создает новый экземпляр тестового репозитория
func NewTestRepository(logger *zap.Logger, db *gorm.DB) TestRepository {
	return &testRepository{
		logger: logger,
		db:     db,
	}
}

// Ping проверяет соединение с БД
func (r *testRepository) Ping(ctx context.Context) error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
