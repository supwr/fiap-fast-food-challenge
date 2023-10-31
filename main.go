package main

import (
	server "github.com/supwr/fiap-fast-food-challenge/src/application/api"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/config"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/database/migration"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log/slog"
)

type AppArgs struct {
	fx.In

	Cfg             config.Config
	Logger          *slog.Logger
	Db              *gorm.DB
	CustomerService *service.CustomerService
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
		migrate(args.Db, args.Cfg, args.Logger)
	}

	api := server.NewApp(args.CustomerService, args.Logger)
	api.Run()
}

func migrate(db *gorm.DB, cfg config.Config, logger *slog.Logger) {
	m := migration.NewMigration(db, cfg, logger)
	m.CreateSchema()
	m.Migrate()
}
