package repository

import "github.com/w-tyler/mcp-test/internal/domain/entity"

type ProductRepository interface {
	GetByID(id string) (*entity.Product, error)
	List() ([]*entity.Product, error)
	Update(product *entity.Product) error
}
