package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/service"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
	"github.com/supwr/fiap-fast-food-challenge/src/infra/http/dto"
	"log/slog"
	"net/http"
	"strconv"
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

func (i *ItemController) ListItems(ctx *gin.Context) {
	var body []dto.Item

	items, err := i.itemService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	jsonItems, _ := json.Marshal(items)
	if err = json.Unmarshal(jsonItems, &body); err != nil {
		i.logger.Error("error enconding item", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, body)
}

func (i *ItemController) GetItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	item, err := i.itemService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if item == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, item)
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

	price := decimal.NewFromFloat(body.Price)
	item := &entity.Item{
		Name:        body.Name,
		Description: body.Description,
		Type:        itemType,
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

func (i *ItemController) DeleteItem(ctx *gin.Context) {
	var err error

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	if err = i.itemService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (i *ItemController) PatchItem(ctx *gin.Context) {
	var item *entity.Item
	var body dto.Item
	var err error

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	if err = ctx.BindJSON(&body); err != nil {
		i.logger.Error("error reading body", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	item, err = body.ToEntity()
	if err != nil {
		i.logger.Error("error getting item", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err = i.itemService.Update(id, item); err != nil {
		i.logger.Error("error updating item", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, item)
}
