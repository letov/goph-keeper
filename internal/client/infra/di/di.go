package di

import (
	"GophKeeper/internal/client/app/aes"
	"GophKeeper/internal/client/app/handler"
	"GophKeeper/internal/client/app/repo"
	"GophKeeper/internal/client/app/session"
	"GophKeeper/internal/client/app/view"
	"GophKeeper/internal/client/infra/config"
	"GophKeeper/internal/client/infra/db"
	"GophKeeper/internal/client/infra/grpcclient"
	"GophKeeper/internal/client/infra/store"
	"GophKeeper/internal/server/infra/logger"

	"go.uber.org/fx"
)

func InjectApp() fx.Option {
	return fx.Provide(
		db.NewDB,
		logger.NewLogger,
		config.NewConfig,
		aes.NewAes,
		session.NewSession,

		handler.NewLogin,
		handler.NewRegister,
		handler.NewSelectDB,

		store.NewRepoDB,
		func(r *store.RepoDB) repo.Snapshot {
			return r
		},

		view.NewRoot,
		grpcclient.NewGrpcClient,
	)
}
