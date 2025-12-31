package model

type Order struct {
	ID        string      `json:"id"`
	Cliente   string      `json:"cliente"`
	Pedido    string      `json:"pedido"`
	Status    OrderStatus `json:"status"`
	CreatedAt string      `json:"created_at"`
}

type OrderStatus string

// 2. Define as opções disponíveis como constantes
const (
	RECEIVED   OrderStatus = "RECEBIDO"
	PROCESSING OrderStatus = "EM_PROCESSAMENTO"
	READY      OrderStatus = "PRONTO"
	FAILED     OrderStatus = "FAILED"
)
