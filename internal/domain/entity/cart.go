package entity

import "github.com/w-tyler/mcp-test/internal/domain/valueobject"

type Cart struct {
	ID      string
	UserID  string
	Items   []CartItem
}

type CartItem struct {
	ProductID string
	Quantity  int
}
