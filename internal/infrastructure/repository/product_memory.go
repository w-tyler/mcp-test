package repository

import (
	"errors"
	"github.com/w-tyler/mcp-test/internal/domain/entity"
)

type ProductMemoryRepo struct {
	products map[string]*entity.Product
}

func NewProductMemoryRepo() *ProductMemoryRepo {
	return &ProductMemoryRepo{products: make(map[string]*entity.Product)}
}

func (r *ProductMemoryRepo) GetByID(id string) (*entity.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (r *ProductMemoryRepo) List() ([]*entity.Product, error) {
	var list []*entity.Product
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

func (r *ProductMemoryRepo) Update(product *entity.Product) error {
	r.products[product.ID] = product
	return nil
}
