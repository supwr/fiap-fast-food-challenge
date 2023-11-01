package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/govalues/decimal"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/dto"
	"log/slog"
	"net/http"
)

type ItemController struct {
	itemService *service.ItemService
	logger      *slog.Logger
}

func NewItemController(i *service.ItemService, l *slog.Logger) *ItemController {
	return &ItemController{
		itemService: i,
		logger:      l,
	}
}

func (i *ItemController) CreateItem(ctx *gin.Context) {
	var body dto.Item
	var err error

	if err = ctx.BindJSON(&body); err != nil {
		i.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	validation := validate(body).Errors
	if len(validation) > 0 {
		i.logger.Error("invalid payload", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, validation)
		return
	}

	itemType, err := valueobject.NewItemType(body.Type)
	if err != nil {
		i.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid item type: %s. allowed types (%s, %s, %s, %s, %s)",
				body.Type,
				valueobject.Beverage,
				valueobject.Food,
				valueobject.Dessert,
				valueobject.Ingredient,
				valueobject.Sides),
		})
		return
	}

	price, err := decimal.NewFromFloat64(body.Price)
	if err != nil {
		i.logger.Error("error getting item price", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating item",
		})
		return
	}

	item := &entity.Item{
		Name:        body.Name,
		Description: body.Description,
		Type:        *itemType,
		Price:       price,
		Active:      true,
	}

	if err = i.itemService.Create(item); err != nil {
		i.logger.Error("error creating item", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating item",
		})
		return
	}

	ctx.JSON(http.StatusOK, body)
}
