package rbac_model

import (
	"context"
	"database/sql"
	"time"

	rbac_domain "github.com/bcdxn/garden-project/internal/domain/rbac"
	"github.com/bcdxn/garden-project/internal/domain/value_type"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Model is the concrete type that implements the rbac repository interface
type Model struct {
	DB *sql.DB
}

// ListRoles retrieves all roles from the database.
func (m *Model) ListRoles(ctx context.Context) ([]rbac_domain.Role, error) {
	type roleDTO struct {
		ID        string       `sql:"id"`
		Name      string       `sql:"name"`
		CreatedAt time.Time    `sql:"created_at"`
		UpdatedAt sql.NullTime `sql:"updated_at"`
	}

	roles := make([]rbac_domain.Role, 0)

	query := "SELECT id, name, created_at, updated_at FROM rbac_role ORDER BY name"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return roles, err
	}

	for rows.Next() {
		var roleRow roleDTO
		err := rows.Scan(&roleRow.ID, &roleRow.Name, &roleRow.CreatedAt, &roleRow.UpdatedAt)
		if err != nil {
			return roles, err
		}

		role := rbac_domain.Role{
			ID:        roleRow.ID,
			Name:      roleRow.Name,
			CreatedAt: roleRow.CreatedAt,
			UpdatedAt: value_type.NullableTime{
				Time:   roleRow.UpdatedAt.Time,
				IsNull: !roleRow.UpdatedAt.Valid,
			},
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// ListPermissonsByRoleID retrieves a list of action/resource (i.e. a permission) pairs that a
// particular role is entitled to.
func (m *Model) ListPermissionsByRoleID(ctx context.Context, id string) ([]rbac_domain.Permission, error) {
	type permissionDTO struct {
		ActionID     string       `sql:"acton.id"`
		ActionName   string       `sql:"action.name"`
		ResourceID   string       `sql:"resource.id"`
		ResourceName string       `sql:"resource.name"`
		CreatedAt    time.Time    `sql:"created_at"`
		UpdatedAt    sql.NullTime `sql:"updated_at"`
	}

	query := `
		SELECT act.id, act.name, res.id, res.name, perm.created_at, perm.updated_at
		FROM rbac_role AS role
		JOIN rbac_permission AS perm ON perm.role_id = role.id
		JOIN rbac_action AS act ON perm.action_id = act.id
		JOIN rbac_resource AS res ON perm.resource_id = res.id
		where role.id = $1
		ORDER BY act.name, res.name
	`

	permissions := make([]rbac_domain.Permission, 0)

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		var permRow permissionDTO
		err := rows.Scan(
			&permRow.ActionID,
			&permRow.ActionName,
			&permRow.ResourceID,
			&permRow.ResourceName,
			&permRow.CreatedAt,
			&permRow.UpdatedAt,
		)

		if err != nil {
			return permissions, err
		}
		perm := rbac_domain.Permission{
			Action:    permRow.ActionName,
			Resource:  permRow.ResourceName,
			CreatedAt: permRow.CreatedAt,
			UpdatedAt: value_type.NullableTime{
				Time:   permRow.UpdatedAt.Time,
				IsNull: !permRow.UpdatedAt.Valid,
			},
		}
		permissions = append(permissions, perm)
	}

	return permissions, nil
}
