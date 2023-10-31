package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/dto"
	"log/slog"
	"net/http"
	"strconv"
)

type CustomerController struct {
	customerService *service.CustomerService
	logger          *slog.Logger
}

func NewCustomerController(c *service.CustomerService, l *slog.Logger) *CustomerController {
	return &CustomerController{
		customerService: c,
		logger:          l,
	}
}

func (c *CustomerController) GetCustomerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	customer, err := c.customerService.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if customer == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var body dto.Customer

	if err := ctx.BindJSON(&body); err != nil {
		c.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, body)
}
