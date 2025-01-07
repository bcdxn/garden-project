package user_model

import (
	"context"
	"database/sql"
	"time"

	user_domain "github.com/bcdxn/garden-project/internal/domain/user"
	"github.com/bcdxn/garden-project/internal/domain/value_type"
)

type Model struct {
	DB *sql.DB
}

func (m *Model) ListUsers(ctx context.Context) ([]user_domain.User, error) {
	type userDTO struct {
		ID          string       `sql:"id"`
		Email       string       `sql:"email"`
		IsVerified  bool         `sql:"is_verified"`
		CreatedAt   time.Time    `sql:"created_at"`
		UpdatedAt   sql.NullTime `sql:"updated_at"`
		LastLoginAt sql.NullTime `sql:"last_login_at"`
	}

	users := make([]user_domain.User, 0)

	query := "SELECT id, email, is_verified, created_at, updated_at, last_login_at FROM app_user ORDER BY id"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var userRow userDTO

		err := rows.Scan(
			&userRow.ID,
			&userRow.Email,
			&userRow.IsVerified,
			&userRow.CreatedAt,
			&userRow.UpdatedAt,
			&userRow.LastLoginAt,
		)
		if err != nil {
			return users, err
		}

		user := user_domain.User{
			ID:         userRow.ID,
			Email:      userRow.Email,
			IsVerified: userRow.IsVerified,
			CreatedAt:  userRow.CreatedAt,
			UpdatedAt: value_type.NullableTime{
				Time:   userRow.UpdatedAt.Time,
				IsNull: !userRow.UpdatedAt.Valid,
			},
			LastLoginAt: value_type.NullableTime{
				Time:   userRow.LastLoginAt.Time,
				IsNull: !userRow.LastLoginAt.Valid,
			},
		}

		users = append(users, user)
	}

	return users, nil
}
