package service

import (
	"github.com/supwr/fiap-fast-food-challenge/src/domain/contract"
	"github.com/supwr/fiap-fast-food-challenge/src/domain/entity"
	"log/slog"
)

type ItemService struct {
	repository contract.ItemRepository
	logger     *slog.Logger
}

func NewItemService(r contract.ItemRepository, l *slog.Logger) *ItemService {
	return &ItemService{
		repository: r,
		logger:     l,
	}
}

func (s *ItemService) Create(i *entity.Item) error {
	return s.repository.Create(i)
}

func (s *ItemService) GetById(id int) (*entity.Item, error) {
	return s.repository.GetById(id)
}
