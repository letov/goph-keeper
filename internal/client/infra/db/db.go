package db

import (
	"GophKeeper/internal/client/infra/config"
	"context"
	"database/sql"
	"embed"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DB struct {
	log zap.SugaredLogger
	mu  sync.Mutex
	db  *sql.DB
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (db *DB) GetDB() *sql.DB {
	return db.db
}

func (db *DB) makeMigrations() {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.db, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Version(db.db, "migrations"); err != nil {
		db.log.Fatal(err)
	}
}

func NewDB(lc fx.Lifecycle, log zap.SugaredLogger, c config.Config) *DB {
	db := &DB{
		log: log,
	}

	if len(c.DatabaseFile) == 0 {
		log.Info("Empty DB file")
		return db
	}

	log.Info("Init DB")

	sqliteDb, err := sql.Open("sqlite3", c.DatabaseFile)
	if err != nil {
		log.Info("Failed init db")
		return db
	}

	db.db = sqliteDb
	db.makeMigrations()

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info("Close DB")
			return db.db.Close()
		},
	})

	return db
}
