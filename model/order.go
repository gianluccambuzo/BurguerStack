package model

type Order struct {
	ID        string `json:"id"`
	Cliente   string `json:"cliente"`
	Pedido    string `json:"pedido"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
