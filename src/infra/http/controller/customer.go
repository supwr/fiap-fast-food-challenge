package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
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

func (c *CustomerController) GetCustomer(ctx *gin.Context) {
	if len(ctx.Param("id")) > 0 {
		c.GetCustomerById(ctx)
		return
	}

	c.GetCustomerByDocument(ctx)
	return
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

func (c *CustomerController) GetCustomerByDocument(ctx *gin.Context) {
	document, err := valueobject.NewDocument(ctx.Query("document"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	customer, err := c.customerService.GetCustomerByDocument(document)

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
	var err error

	if err = ctx.BindJSON(&body); err != nil {
		c.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	validation := validate(body).Errors
	if len(validation) > 0 {
		c.logger.Error("invalid payload", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, validation)
		return
	}

	document, err := valueobject.NewDocument(body.Document)
	if err != nil {
		c.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid document",
		})
		return
	}

	customer := &entity.Customer{
		Name:     body.Name,
		Document: *document,
		Email:    body.Email,
	}

	if err = c.customerService.Create(customer); err != nil {
		c.logger.Error("error creating customer", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, body)
}
