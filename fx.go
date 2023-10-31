package main

import (
	server "github.com/supwr/fiap-fast-food-challenge/src/application/api"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/contract"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/config"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/database/migration"
	database "github.com/supwr/fiap-fast-food-challenge/src/infra/database/postgres"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func createApp(o ...fx.Option) *fx.App {
	options := []fx.Option{
		fx.Provide(
			newConfig,
			newLogger,
			newConnection,
			newMigration,
			newApp,

			// repositories
			fx.Annotate(
				repository.NewCustomerRepository,
				fx.As(new(contract.CustomerRepository)),
			),

			//services
			service.NewCustomerService,
		),
	}

	return fx.New(append(options, o...)...)
}

func newConfig() (config.Config, error) {
	return config.NewConfig()
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}

func newConnection(cfg config.Config) (*gorm.DB, error) {
	return database.NewConnection(
		cfg,
	)
}

func newMigration(db *gorm.DB, cfg config.Config, logger *slog.Logger) *migration.Migration {
	return migration.NewMigration(db, cfg, logger)
}

func newApp(args server.AppArgs) *server.App {
	return server.NewApp(args)
}
