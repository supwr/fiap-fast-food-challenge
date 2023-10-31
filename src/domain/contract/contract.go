package contract

import "github.com/supwr/fiap-fast-food-challenge/src/domain/entity"

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	GetById(id int) (*entity.Customer, error)
}
