package repository

import (
	"github.com/txzy2/go-logger-api/pkg/database"
)

type Repository struct {
	TestRepository TestRepository
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		TestRepository: NewTestRepository(db.GORM),
	}
}
