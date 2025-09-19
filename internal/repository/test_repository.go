package repository

import (
	"context"
	"database/sql"
)

type TestRepository interface {
	Ping(ctx context.Context) error
}

type testRepository struct {
	db *sql.DB
}

// NewTestRepository создает новый экземпляр тестового репозитория
func NewTestRepository(db *sql.DB) TestRepository {
	return &testRepository{
		db: db,
	}
}

// Ping проверяет соединение с БД
func (r *testRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}
