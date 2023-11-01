package dto

import (
	"github.com/govalues/decimal"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
)

type Item struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"type,omitempty" validate:"required"`
	Price       float64 `json:"price,omitempty" required:"required"`
}

func (i *Item) ToEntity() (*entity.Item, error) {
	var iType *valueobject.ItemType
	var err error

	if len(i.Type) > 0 {
		iType, err = valueobject.NewItemType(i.Type)
		if err != nil {
			return nil, err
		}
	}

	price, err := decimal.NewFromFloat64(i.Price)
	if err != nil {
		return nil, err
	}

	return &entity.Item{
		Name:        i.Name,
		Description: i.Description,
		Type:        iType,
		Price:       price,
	}, nil
}
