package aggregate

import "github.com/w-tyler/mcp-test/internal/domain/entity"

type CartAggregate struct {
	Cart entity.Cart
}

func (ca *CartAggregate) AddItem(item entity.CartItem) {
	ca.Cart.Items = append(ca.Cart.Items, item)
}

func (ca *CartAggregate) RemoveItem(productID string) {
	items := ca.Cart.Items
	for i, item := range items {
		if item.ProductID == productID {
			ca.Cart.Items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
