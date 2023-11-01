package repository

import (
	"errors"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"gorm.io/gorm"
	"log/slog"
)

type ItemRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewItemRepository(db *gorm.DB, logger *slog.Logger) *ItemRepository {
	return &ItemRepository{
		db:     db,
		logger: logger,
	}
}

func (i *ItemRepository) Create(item *entity.Item) error {
	return i.db.Create(item).Error
}

func (i *ItemRepository) GetById(id int) (*entity.Item, error) {
	var item *entity.Item

	if err := i.db.First(&item, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		i.logger.Error("error finding item", slog.Any("error", err))
		return nil, err
	}

	return item, nil
}
