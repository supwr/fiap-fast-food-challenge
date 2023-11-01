package entity

import (
	"github.com/govalues/decimal"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
	"time"
)

type Item struct {
	ID          uint                  `json:"id" gorm:"primaryKey"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Type        *valueobject.ItemType `json:"type"`
	Price       decimal.Decimal       `json:"price"`
	Active      bool                  `json:"active"`
	CreatedAt   *time.Time            `json:"created_at"`
	UpdatedAt   *time.Time            `json:"updated_at"`
	DeletedAt   *time.Time            `json:"deleted_at"`
}
