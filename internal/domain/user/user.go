package user_domain

import (
	"context"
	"time"

	"github.com/bcdxn/garden-project/internal/domain/value_type"
)

type User struct {
	ID          string
	Email       string
	IsVerified  bool
	CreatedAt   time.Time
	UpdatedAt   value_type.NullableTime
	LastLoginAt value_type.NullableTime
}

type Repository interface {
	ListUsers(ctx context.Context) ([]User, error)
}
