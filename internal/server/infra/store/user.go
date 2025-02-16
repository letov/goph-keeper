package store

import (
	"GophKeeper/internal/common/dto"
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (r *RepoDB) SaveUser(ctx context.Context, u dto.SaveUser) error {
	query := `INSERT INTO users (email, pass_hash) VALUES (@email, @pass_hash)`
	args := pgx.NamedArgs{
		"email":     u.Email,
		"pass_hash": u.PassHash,
	}

	_, err := r.pool.Exec(ctx, query, args)
	return err
}

func (r *RepoDB) LoginUser(ctx context.Context, l dto.LoginUser) (int32, error) {
	query := `SELECT id FROM users WHERE LOWER(email)=@email AND pass_hash=@pass_hash`
	args := pgx.NamedArgs{
		"email":     strings.ToLower(l.Email),
		"pass_hash": l.PassHash,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int32
	rows.Next()
	err = rows.Scan(&id)
	return id, err
}
