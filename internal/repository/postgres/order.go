package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type Order struct {
	conn  *sql.DB
	table string
}

const (
	tableOrder = "order"
)

func NewOrder(conn *sql.DB) *Order {
	return &Order{
		conn:  conn,
		table: tableOrder,
	}
}

func (r *Order) Create(ctx context.Context, order model.Order) error {
	object := dao.FromOrder(order)
	query := "INSERT INTO " + r.table + " (customer_name, status, notes) VALUES ($1, $2, $3)"

	_, err := r.conn.Exec(query, object.CustomerName, object.Status, object.Notes)
	if err != nil {
		return err
	}

	return nil
}

func (r *Order) Get(ctx context.Context, id int) (model.Order, error) {
	var order dao.Order
	query := "SELECT order_id, customer_name, status, notes, created_at FROM " + r.table + " WHERE id = $1"

	err := r.conn.QueryRow(query, id).Scan(&order.OrderID, &order.CustomerName, &order.Status, &order.Notes, &order.CreatedAt)
	if err != nil {
		return model.Order{}, nil
	}

	return dao.ToOrder(order), nil
}

func (r *Order) GetAll(ctx context.Context) ([]model.Order, error) {
	var order_all []model.Order
	query := "SELECT * FROM " + r.table

	rows, err := r.conn.Query(query)
	if err != nil {
		return []model.Order{}, nil
	}

	for rows.Next() {
		var order dao.Order
		err := rows.Scan(&order.OrderID, &order.CustomerName, &order.Status, &order.Notes, &order.CreatedAt)
		if err != nil {
			return []model.Order{}, nil
		}

		order_all = append(order_all, dao.ToOrder(order))
	}

	return order_all, nil
}

func (r *Order) UPDATE(ctx context.Context, order model.Order, id int) error {
	object := dao.FromOrder(order)
	query := "UPDATE " + r.table + " SET customer_name = $1, status = $2, notes = $3 WHERE id = $4"

	_, err := r.conn.Exec(query, object.CustomerName, object.Status, object.Notes, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Order) DELETE(ctx context.Context, id int) error {
	query := "DELETE FROM " + r.table + " WHERE id = $1"

	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
