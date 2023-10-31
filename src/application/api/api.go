package api

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/controller"
	"log/slog"
)

type App struct {
	customerService *service.CustomerService
	logger          *slog.Logger
}

func NewApp(c *service.CustomerService, l *slog.Logger) *App {
	return &App{
		customerService: c,
		logger:          l,
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
