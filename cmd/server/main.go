package main

import (
	"GophKeeper/internal/server/app/app"
	"GophKeeper/internal/server/infra/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(app.Start),
	).Run()
}
