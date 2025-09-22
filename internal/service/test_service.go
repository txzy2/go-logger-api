package service

import (
	"context"

	"github.com/txzy2/go-logger-api/internal/repository"
	"go.uber.org/zap"
)

type TestService interface {
	Ping(ctx context.Context) error
}

type testService struct {
	logger   *zap.Logger
	testRepo repository.TestRepository
}

func NewTestService(logger *zap.Logger, testRepo repository.TestRepository) TestService {
	return &testService{
		logger:   logger,
		testRepo: testRepo,
	}
}

func (s *testService) Ping(ctx context.Context) error {
	return s.testRepo.Ping(ctx)
}
