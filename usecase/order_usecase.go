package usecase

import (
	"example/BurgerStack/model"
	"example/BurgerStack/repository"
	"github.com/google/uuid"
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

func (ou *OrderUseCase) GetOrderById(orderId string) (*model.Order, error) {
	return ou.OrderRepository.GetOrderById(orderId)
}

func (ou *OrderUseCase) CreateOrder(order model.Order) (model.Order, error) {
	order.ID = uuid.NewString()
	order.Status = "RECEBIDO"

	return ou.OrderRepository.InsertOrder(order)
}
