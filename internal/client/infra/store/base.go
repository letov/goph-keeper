package store

import (
	"GophKeeper/internal/client/infra/db"
	"database/sql"

	"go.uber.org/zap"
)

type RepoDB struct {
	db  *sql.DB
	log zap.SugaredLogger
}

func NewRepoDB(db *db.DB, log zap.SugaredLogger) *RepoDB {
	return &RepoDB{
		db:  db.GetDB(),
		log: log,
	}
}
