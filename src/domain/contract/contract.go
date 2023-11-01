package contract

import (
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	GetById(id int) (*entity.Customer, error)
	GetCustomerByDocument(document *valueobject.Document) (*entity.Customer, error)
}

type ItemRepository interface {
	Create(item *entity.Item) error
	GetById(id int) (*entity.Item, error)
}
