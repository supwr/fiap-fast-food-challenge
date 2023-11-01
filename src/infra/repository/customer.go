package repository

import (
	"errors"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/valueobject"
	"gorm.io/gorm"
	"log/slog"
)

type CustomerRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewCustomerRepository(db *gorm.DB, logger *slog.Logger) *CustomerRepository {
	return &CustomerRepository{
		db:     db,
		logger: logger,
	}
}

func (c *CustomerRepository) Create(customer *entity.Customer) error {
	return c.db.Create(customer).Error
}

func (c *CustomerRepository) GetById(id int) (*entity.Customer, error) {
	var customer *entity.Customer

	if err := c.db.First(&customer, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("error finding customer", slog.Any("error", err))
		return nil, err
	}

	return customer, nil
}

func (c *CustomerRepository) GetCustomerByDocument(document *valueobject.Document) (*entity.Customer, error) {
	var customer *entity.Customer

	if err := c.db.First(&customer, "document = ?", document.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("error finding customer", slog.Any("error", err))
		return nil, err
	}

	return customer, nil
}
