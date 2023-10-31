package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"net/http"
	"strconv"
)

type CustomerController struct {
	customerService *service.CustomerService
}

func NewCustomerController(c *service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: c,
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
