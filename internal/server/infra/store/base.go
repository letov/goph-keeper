package store

import (
	"GophKeeper/internal/server/infra/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type RepoDB struct {
	pool *pgxpool.Pool
	log  zap.SugaredLogger
}

func NewRepoDB(db *db.DB, log zap.SugaredLogger) *RepoDB {
	return &RepoDB{
		db.GetPool(),
		log,
	}
}
