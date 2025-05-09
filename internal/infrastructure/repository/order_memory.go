package repository

import (
	"errors"
	"github.com/w-tyler/mcp-test/internal/domain/entity"
)

type OrderMemoryRepo struct {
	orders map[string]*entity.Order
}

func NewOrderMemoryRepo() *OrderMemoryRepo {
	return &OrderMemoryRepo{orders: make(map[string]*entity.Order)}
}

func (r *OrderMemoryRepo) Save(order *entity.Order) error {
	r.orders[order.ID] = order
	return nil
}

func (r *OrderMemoryRepo) GetByID(id string) (*entity.Order, error) {
	o, ok := r.orders[id]
	if !ok {
		return nil, errors.New("order not found")
	}
	return o, nil
}

func (r *OrderMemoryRepo) List() ([]*entity.Order, error) {
	var list []*entity.Order
	for _, o := range r.orders {
		list = append(list, o)
	}
	return list, nil
}
