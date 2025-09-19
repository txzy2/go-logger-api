package service

import (
	"context"

	"github.com/txzy2/go-logger-api/internal/repository"
)

type TestService interface {
	Ping(ctx context.Context) error
}

type testService struct {
	testRepo repository.TestRepository
}

func NewTestService(testRepo repository.TestRepository) TestService {
	return &testService{
		testRepo: testRepo,
	}
}

func (s *testService) Ping(ctx context.Context) error {
	return s.testRepo.Ping(ctx)
}
