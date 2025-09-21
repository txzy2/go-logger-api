package repository

import (
	"context"

	"gorm.io/gorm"
)

type TestRepository interface {
	Ping(ctx context.Context) error
}

type testRepository struct {
	db *gorm.DB
}

// NewTestRepository создает новый экземпляр тестового репозитория
func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{
		db: db,
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
