package service

import "github.com/w-tyler/mcp-test/internal/domain/entity"

type InventoryService struct{}

func (is *InventoryService) CheckStock(product *entity.Product, quantity int) bool {
	return product.Stock >= quantity
}

func (is *InventoryService) DecreaseStock(product *entity.Product, quantity int) {
	if product.Stock >= quantity {
		product.Stock -= quantity
	}
}
