package contract

import "github.com/supwr/fiap-fast-food-challenge/src/domain/entity"

type CustomerReposity interface {
	Create(customer *entity.Customer) error
	GetById(id int) (*entity.Customer, error)
}
