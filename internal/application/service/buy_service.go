package service

import (
	"github.com/w-tyler/mcp-test/internal/domain/entity"
	"github.com/w-tyler/mcp-test/internal/domain/valueobject"
	"github.com/w-tyler/mcp-test/internal/domain/repository"
	"github.com/google/uuid"
)

type BuyService struct {
	UserRepo    repository.UserRepository
	ProductRepo repository.ProductRepository
	OrderRepo   repository.OrderRepository
}

func (s *BuyService) Buy(userID, productID string, quantity int) (*entity.Order, error) {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	product, err := s.ProductRepo.GetByID(productID)
	if err != nil {
		return nil, err
	}
	if product.Stock < quantity {
		return nil, errors.New("not enough stock")
	}
	product.Stock -= quantity
	s.ProductRepo.Update(product)

	orderItem := entity.OrderItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     valueobject.Money{Amount: product.Price.Amount, Currency: product.Price.Currency},
	}
	order := &entity.Order{
		ID:      uuid.New().String(),
		UserID:  userID,
		Items:   []entity.OrderItem{orderItem},
		Total:   valueobject.Money{Amount: product.Price.Amount * float64(quantity), Currency: product.Price.Currency},
		Status:  valueobject.OrderPending,
		Address: user.Address,
	}
	s.OrderRepo.Save(order)
	return order, nil
}
