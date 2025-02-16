package main

import (
	"GophKeeper/internal/client/app/app"
	"GophKeeper/internal/client/infra/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(app.Start),
	).Run()
}
