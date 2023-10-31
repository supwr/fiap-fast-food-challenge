package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/config"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/database/migration"
	database "github.com/supwr/fiap-fast-food-challenge/src/infra/database/postgres"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/controller"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/repository"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	var err error
	var logger = loadLogger()

	r := gin.Default()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := loadDatabase(cfg)

	if err != nil {
		logger.Error("error connecting to database", err)
		panic(err)
	}

	if cfg.Environment == "DEV" {
		migrate(db, cfg, logger)
	}

	customerRepository := repository.NewCustomerRepository(db, logger)
	customerService := service.NewCustomerService(customerRepository)

	customerController := controller.NewCustomerController(customerService)

	r.GET("/customers/:id", customerController.GetCustomerById)

	r.Run()
}

func migrate(db *gorm.DB, cfg config.Config, logger *slog.Logger) {
	m := migration.NewMigration(db, cfg, logger)
	m.CreateSchema()
	m.Migrate()
}

func loadDatabase(cfg config.Config) (*gorm.DB, error) {
	return database.NewConnection(
		cfg,
	)
}

func loadConfig() (config.Config, error) {
	return config.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
