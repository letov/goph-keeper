package di

import (
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/config"
	"GophKeeper/internal/server/infra/db"
	"GophKeeper/internal/server/infra/grpcserver"
	"GophKeeper/internal/server/infra/logger"
	"GophKeeper/internal/server/infra/store"

	"go.uber.org/fx"
)

func InjectApp() fx.Option {
	return fx.Provide(
		config.NewConfig,
		logger.NewLogger,
		db.NewDB,

		store.NewUserRepoDB,

		func(ur *store.UserRepoDB) repo.User {
			return ur
		},

		grpcserver.NewGrpcServer,
	)
}
