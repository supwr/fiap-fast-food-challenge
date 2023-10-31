package service

import (
	"github.com/supwr/fiap-fast-food-challenge/src/domain/contract"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
)

type CustomerService struct {
	repository contract.CustomerReposity
}

func NewCustomerService(r contract.CustomerReposity) *CustomerService {
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
