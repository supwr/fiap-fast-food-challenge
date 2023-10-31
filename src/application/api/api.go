package api

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/controller"
	"go.uber.org/fx"
	"log/slog"
)

type AppArgs struct {
	fx.In

	CustomerService *service.CustomerService
	Logger          *slog.Logger
}

type App struct {
	customerService *service.CustomerService
	logger          *slog.Logger
}

func NewApp(a AppArgs) *App {
	return &App{
		customerService: a.CustomerService,
		logger:          a.Logger,
	}
}

func (a *App) Run() {
	app := gin.Default()

	//controllers
	customerController := controller.NewCustomerController(a.customerService, a.logger)

	// routes
	app.GET("/customers/:id", customerController.GetCustomerById)
	app.POST("/customers", customerController.CreateCustomer)

	// app run
	app.Run()
}
