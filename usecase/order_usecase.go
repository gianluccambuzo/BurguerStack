package usecase

import (
	"example/BurgerStack/model"
	"example/BurgerStack/repository"
)

type OrderUseCase struct {
	OrderRepository repository.OrderRepository
}

func NewOrderUseCase(orderRepository repository.OrderRepository) OrderUseCase {
	return OrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (ou *OrderUseCase) GetOrders() ([]model.Order, error) {
	return ou.OrderRepository.GetOrders()
}
