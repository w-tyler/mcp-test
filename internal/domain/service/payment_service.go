package service

import (
	"github.com/w-tyler/mcp-test/internal/domain/entity"
	"github.com/w-tyler/mcp-test/internal/domain/valueobject"
)

type PaymentService struct{}

func (ps *PaymentService) Pay(order *entity.Order, amount valueobject.Money) bool {
	// Dummy payment logic
	if amount.Amount >= order.Total.Amount {
		order.Status = valueobject.OrderPaid
		return true
	}
	return false
}
