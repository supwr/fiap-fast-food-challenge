package repository

import (
	"errors"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"gorm.io/gorm"
	"log/slog"
	"time"
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

func (i *ItemRepository) Update(id int, item *entity.Item) error {
	return i.db.Model(&item).Where("id = ? and deleted_at is null", id).Updates(item).Error
}

func (i *ItemRepository) Delete(id int) error {
	var item entity.Item
	return i.db.Model(&item).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (i *ItemRepository) List() ([]*entity.Item, error) {
	var items []*entity.Item

	if err := i.db.Find(&items, "deleted_at is null").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		i.logger.Error("error listing items", slog.Any("error", err))
		return nil, err
	}

	return items, nil
}

func (i *ItemRepository) GetById(id int) (*entity.Item, error) {
	var item *entity.Item

	if err := i.db.First(&item, "id = ? and deleted_at is null", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		i.logger.Error("error finding item", slog.Any("error", err))
		return nil, err
	}

	return item, nil
}
