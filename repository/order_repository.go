package repository

import (
	"database/sql"
	"example/BurgerStack/model"
	"fmt"
)

type OrderRepository struct {
	connection *sql.DB
}

func NewOrderRepository(connection *sql.DB) OrderRepository {
	return OrderRepository{connection: connection}
}

func (ou *OrderRepository) GetOrders() ([]model.Order, error) {

	query := "SELECT * FROM pedidos ORDER BY id DESC"
	rows, err := ou.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Order{}, err
	}

	var orders []model.Order
	var order model.Order

	for rows.Next() {
		err = rows.Scan(
			&order.ID,
			&order.Cliente,
			&order.Pedido,
			&order.Status,
			&order.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.Order{}, err
		}

		orders = append(orders, order)
	}

	rows.Close()

	return orders, nil

}
