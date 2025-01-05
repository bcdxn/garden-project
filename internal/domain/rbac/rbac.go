package rbac_domain

import (
	"context"
	"time"

	"github.com/bcdxn/garden-project/internal/domain/value_type"
)

type Role struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt value_type.NullableTime
}

type Action struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Resource struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	Action    Action
	Resource  Resource
	CreatedAt time.Time
	UpdatedAt value_type.NullableTime
}

type Repository interface {
	ListRoles(ctx context.Context) ([]Role, error)
	ListPermissionsByRoleID(ctx context.Context, roleId string) ([]Permission, error)
}
