package rbac_domain

import "time"

type Role struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
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
	Action   Action
	Resource Resource
}

type Repository interface {
	ListRoles() ([]Role, error)
	ListPermissionsByRoleID(roleId string) ([]Permission, error)
}
