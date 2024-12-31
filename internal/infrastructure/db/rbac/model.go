package rbac_model

import (
	"errors"

	rbac_domain "github.com/bcdxn/garden-project/internal/domain/rbac"
)

// type Repository interface {
// 	ListRoles() ([]Role, error)
// 	ListPermissionsByRoleID(roleId string) ([]Permission, error)
// }

type Model struct {
	// db *sql.DB
}

func (m *Model) ListRoles() ([]rbac_domain.Role, error) {
	var roles []rbac_domain.Role
	return roles, errors.ErrUnsupported
}

func (m *Model) ListPermissionsByRoleID(id string) ([]rbac_domain.Permission, error) {
	var permissions []rbac_domain.Permission
	return permissions, errors.ErrUnsupported
}
