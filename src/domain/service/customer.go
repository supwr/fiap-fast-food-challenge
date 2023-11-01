package service

import (
	"github.com/supwr/fiap-fast-food-challenge/src/domain/contract"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
)

type CustomerService struct {
	repository contract.CustomerRepository
}

func NewCustomerService(r contract.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: r,
	}
}

func (c *CustomerService) Create(customer *entity.Customer) error {
	return c.repository.Create(customer)
}

func (c *CustomerService) GetById(id int) (*entity.Customer, error) {
	return c.repository.GetById(id)
}

func (c *CustomerService) GetCustomerByDocument(document *valueobject.Document) (*entity.Customer, error) {
	return c.repository.GetCustomerByDocument(document)
}
