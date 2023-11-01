package api

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/controller"
	"go.uber.org/fx"
	"log/slog"
)

type AppArgs struct {
	fx.In

	CustomerController *controller.CustomerController
	ItemController     *controller.ItemController
	Logger             *slog.Logger
}

type App struct {
	customerController *controller.CustomerController
	itemController     *controller.ItemController
	logger             *slog.Logger
}

func NewApp(a AppArgs) *App {
	return &App{
		customerController: a.CustomerController,
		itemController:     a.ItemController,
		logger:             a.Logger,
	}
}

func (a *App) Run() {
	app := gin.Default()

	// routes
	app.GET("/customers", a.customerController.GetCustomerByDocument)
	app.GET("/customers/:id", a.customerController.GetCustomerById)
	app.POST("/customers", a.customerController.CreateCustomer)

	app.POST("/items", a.itemController.CreateItem)

	// app run
	app.Run()
}
