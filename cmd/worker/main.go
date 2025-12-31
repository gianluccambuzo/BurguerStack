package main

import (
	"context"
	"encoding/json"
	"example/BurgerStack/config/db"
	"example/BurgerStack/repository"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	sqsConfig "example/BurgerStack/config/sqs"
	"example/BurgerStack/model"
	"example/BurgerStack/usecase"
)

func main() {

	sqsClient := sqsConfig.NewSqsClient()
	queueUrl := sqsConfig.GetQueueUrl()

	fmt.Println("Worker iniciado via LocalStack. Aguardando mensagens...")

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	OrderRepository := repository.NewOrderRepository(dbConnection)
	OrderUseCase := usecase.NewOrderUseCase(OrderRepository)

	for {

		output, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueUrl),
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     20,
			VisibilityTimeout:   60,
		})

		if err != nil {
			log.Printf("Error receiving message from queue: %s\n", err)
		}

		//Cria Wait Group
		var wg sync.WaitGroup

		for _, msg := range output.Messages {

			//Adiciona contador ao WaitGroup
			wg.Add(1)

			go func(m types.Message) {

				//Diminui a contagem do WG quando estiver Done
				defer wg.Done()

				err := processMessage(sqsClient, queueUrl, msg, OrderUseCase)
				if err != nil {
					log.Printf("Erro ao processar mensagem %s: %v", *msg.MessageId, err)
				}

			}(msg)

		}

		//Espera o WorkGroup finalizar
		wg.Wait()

		log.Println("Lote finalizado. Buscando próximas...")
	}
}

func processMessage(client *sqs.Client, queueURL string, msg types.Message, useCase usecase.OrderUseCase) error {
	// A. Parse do JSON (Body da mensagem) para o Struct
	var order model.Order

	// O Body é um ponteiro para string (*string), precisamos pegar o valor e converter para bytes
	if msg.Body == nil {
		return fmt.Errorf("corpo da mensagem vazio")
	}

	err := json.Unmarshal([]byte(*msg.Body), &order)
	if err != nil {
		return fmt.Errorf("falha ao decodificar JSON: %w", err)
	}

	// B. Lógica de Negócio (aqui você salvaria no banco, enviaria email, etc)
	fmt.Printf("📦 Processando Pedido: ID=%s | Cliente=%s | Pedido=%s | Status=%s \n", order.ID, order.Cliente, order.Pedido, order.Status)
	time.Sleep(10 * time.Second)
	useCase.UpdateOrderStatus(order.ID, model.PROCESSING)
	fmt.Printf("Pedido processado ID=%s alterado para status %s\n", order.ID, model.PROCESSING)

	time.Sleep(10 * time.Second)
	useCase.UpdateOrderStatus(order.ID, model.READY)
	fmt.Printf("Pedido processado ID=%s alterado para status %s\n", order.ID, model.READY)

	// C. Deletar a mensagem da fila (Passo CRUCIAL)
	// Se não deletar, o SQS entrega ela de novo depois do VisibilityTimeout
	_, err = client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: msg.ReceiptHandle, // O identificador para deletar é o ReceiptHandle, não o ID
	})

	if err != nil {
		return fmt.Errorf("falha ao deletar mensagem do SQS: %w", err)
	}

	return nil
}
