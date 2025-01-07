package user_app

import (
	"context"

	user_domain "github.com/bcdxn/garden-project/internal/domain/user"
)

type Service struct {
	repository user_domain.Repository
}

func NewService(repository user_domain.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListUsers(ctx context.Context) ([]user_domain.User, error) {
	return s.repository.ListUsers(ctx)
}
