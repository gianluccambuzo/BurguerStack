#!/bin/bash
echo "🏗️  Inicializando recursos do LocalStack..."

# Cria a fila SQS
awslocal sqs create-queue --queue-name fila-pedidos

echo "✅ Fila 'fila-pedidos' criada com sucesso!"