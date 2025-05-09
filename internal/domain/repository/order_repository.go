package repository

import "github.com/w-tyler/mcp-test/internal/domain/entity"

type OrderRepository interface {
	Save(order *entity.Order) error
	GetByID(id string) (*entity.Order, error)
	List() ([]*entity.Order, error)
}
