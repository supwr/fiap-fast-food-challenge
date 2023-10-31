package main

import (
	server "github.com/supwr/fiap-fast-food-challenge/src/application/api"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/config"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/database/migration"
	"go.uber.org/fx"
)

type AppArgs struct {
	fx.In

	Cfg       config.Config
	Migration *migration.Migration
	Api       *server.App
}

const devEnv = "DEV"

func main() {
	app := createApp(
		fx.Invoke(
			runAPI,
		),
		fx.Invoke(func(s fx.Shutdowner) { _ = s.Shutdown() }),
	)

	app.Run()
}

func runAPI(args AppArgs) {
	if args.Cfg.Environment == devEnv {
		migrate(args.Migration)
	}

	args.Api.Run()
}

func migrate(m *migration.Migration) {
	m.CreateSchema()
	m.Migrate()
}
