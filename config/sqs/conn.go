package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	host      = "http://sqs.us-east-1.localhost.localstack.cloud:4566/"
	accountId = "000000000000"
	queueName = "fila-pedidos"
)

func NewSqsClient() *sqs.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)

	if err != nil {
		log.Fatalf("Erro ao carregar config AWS: %v", err)
	}

	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		// Esta é a nova forma de apontar para o LocalStack
		// Substitui o antigo EndpointResolverWithOptions
		o.BaseEndpoint = aws.String(host)
	})

	return client
}

func GetQueueUrl() string {
	queue := host + accountId + "/" + queueName
	fmt.Println(queue)
	return queue
}
