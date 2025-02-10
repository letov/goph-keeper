package store

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/internal/server/infra/db"
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type UserRepoDB struct {
	db  *db.DB
	log zap.SugaredLogger
}

func (ur *UserRepoDB) Save(ctx context.Context, u dto.SaveUser) error {
	query := `INSERT INTO users (email, pass_hash, public_key) VALUES (@email, @pass_hash, @public_key)`
	args := pgx.NamedArgs{
		"email":      u.Email,
		"pass_hash":  u.PassHash,
		"public_key": u.PublicKey,
	}

	job := ur.db.ExecJob(query, args)
	_, err := job(ctx)
	return err
}

func (ur *UserRepoDB) Login(ctx context.Context, l dto.LoginUser) (bool, error) {
	query := `SELECT count(id) as count FROM users WHERE email=@email AND pass_hash=@pass_hash`
	args := pgx.NamedArgs{
		"email":     l.Email,
		"pass_hash": l.PassHash,
	}

	job := ur.db.QueryJob(query, args)
	rows, err := job(ctx)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		return count == 1, err
	}
	return false, nil
}

func NewUserRepoDB(db *db.DB, log zap.SugaredLogger) *UserRepoDB {
	return &UserRepoDB{
		db,
		log,
	}
}
