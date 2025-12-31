package usecase

import (
	"context"
	"encoding/json"
	sqsConfig "example/BurgerStack/config/sqs"
	"example/BurgerStack/model"
	"example/BurgerStack/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	"log"
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

	var insertedOrder, err = ou.OrderRepository.InsertOrder(order)

	publishOrderSqs(insertedOrder)

	return insertedOrder, err
}

func (ou *OrderUseCase) UpdateOrderStatus(id string, status string) {

	ou.OrderRepository.UpdateOrderStatus(id, status)
}

func publishOrderSqs(order model.Order) {

	sqsClient := sqsConfig.NewSqsClient()
	queueUrl := sqsConfig.GetQueueUrl()

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Printf("Erro ao fazer marshal do pedido: %v", err)
		return
	}

	_, err = sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: aws.String(string(orderJSON)),
		QueueUrl:    aws.String(queueUrl),
	})

	if err != nil {
		log.Printf("Erro ao enviar para o SQS: %v", err)
		return
	}

	log.Printf("Sucesso!!")
}
