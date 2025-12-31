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

func (ou *OrderRepository) GetOrderById(orderId string) (*model.Order, error) {

	query, err := ou.connection.Prepare("SELECT * FROM pedidos WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var order model.Order
	err = query.QueryRow(orderId).Scan(
		&order.ID, &order.Cliente, &order.Pedido, &order.Status, &order.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &order, nil
}

func (ou *OrderRepository) InsertOrder(order model.Order) (model.Order, error) {

	query, err := ou.connection.Prepare("INSERT INTO pedidos " +
		"(id, cliente, item, status) " +
		"VALUES($1, $2, $3, $4) RETURNING id, cliente, item, status, created_at")

	if err != nil {
		fmt.Println(err)
		return model.Order{}, err
	}

	err = query.QueryRow(order.ID, order.Cliente, order.Pedido, order.Status).Scan(
		&order.ID, &order.Cliente, &order.Pedido, &order.Status, &order.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return model.Order{}, err
	}

	query.Close()
	return order, nil
}

func (ou *OrderRepository) UpdateOrderStatus(id string, status model.OrderStatus) {
	query, err := ou.connection.Prepare("UPDATE pedidos " +
		"SET status = $2 WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return
	}

	query.QueryRow(id, status)
	query.Close()
}
