package rbac_app

import rbac_domain "github.com/bcdxn/garden-project/internal/domain/rbac"

type Service struct {
	repository rbac_domain.Repository
}

func NewService(repository rbac_domain.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// ListRoles retrieves a list of all roles in the repository.
func (s *Service) ListRoles() ([]rbac_domain.Role, error) {
	return s.repository.ListRoles()
}

// ListPermissionsByRoleID retrieves a list of permissions for a specicified role. A 'Permission' is
// defined as the combination of an allowed 'Action' on a specified 'Resource'.
func (s *Service) ListPermissionsByRoleID(id string) ([]rbac_domain.Permission, error) {
	return s.repository.ListPermissionsByRoleID(id)
}
