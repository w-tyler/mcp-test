package aggregate

import (
	"github.com/w-tyler/mcp-test/internal/domain/entity"
	"github.com/w-tyler/mcp-test/internal/domain/valueobject"
)

type OrderAggregate struct {
	Order entity.Order
}

func (oa *OrderAggregate) AddItem(item entity.OrderItem) {
	oa.Order.Items = append(oa.Order.Items, item)
}

func (oa *OrderAggregate) CalculateTotal() valueobject.Money {
	total := 0.0
	currency := "USD"
	for _, item := range oa.Order.Items {
		total += item.Price.Amount * float64(item.Quantity)
		currency = item.Price.Currency
	}
	return valueobject.Money{Amount: total, Currency: currency}
}
