package rbac_model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	rbac_domain "github.com/bcdxn/garden-project/internal/domain/rbac"
	"github.com/bcdxn/garden-project/internal/domain/value_type"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Model struct {
	DB *sql.DB
}

func (m *Model) ListRoles(ctx context.Context) ([]rbac_domain.Role, error) {
	var roles []rbac_domain.Role

	type roleDTO struct {
		ID        string       `sql:"id"`
		Name      string       `sql:"name"`
		CreatedAt time.Time    `sql:"created_at"`
		UpdatedAt sql.NullTime `sql:"updated_at"`
	}

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

func (m *Model) ListPermissionsByRoleID(cts context.Context, id string) ([]rbac_domain.Permission, error) {
	var permissions []rbac_domain.Permission
	return permissions, errors.ErrUnsupported
}
