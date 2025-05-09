package entity

import "github.com/w-tyler/mcp-test/internal/domain/valueobject"

type Order struct {
	ID         string
	UserID     string
	Items      []OrderItem
	Total      valueobject.Money
	Status     valueobject.OrderStatus
	Address    valueobject.Address
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     valueobject.Money
}
