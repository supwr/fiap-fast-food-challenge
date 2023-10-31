package entity

import (
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
	"time"
)

type Customer struct {
	ID        uint                 `json:"id" gorm:"primaryKey"`
	Name      string               `json:"name"`
	Document  valueobject.Document `json:"document"`
	Email     string               `json:"email"`
	CreatedAt *time.Time           `json:"created_at"`
	UpdatedAt *time.Time           `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at"`
}
