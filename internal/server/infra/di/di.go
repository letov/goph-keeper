package di

import (
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/auth"
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

		store.NewRepoDB,

		func(r *store.RepoDB) repo.User {
			return r
		},
		func(r *store.RepoDB) repo.Snapshot {
			return r
		},

		auth.NewService,
		grpcserver.NewAuthInterceptor,
		grpcserver.NewGrpcServer,
	)
}
