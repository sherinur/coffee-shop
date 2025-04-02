package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type OrderItems struct {
	conn  *sql.DB
	table string
}

const (
	tableOrderItems = "order_items"
)

func NewOrderItems(conn *sql.DB) *OrderItems {
	return &OrderItems{
		conn:  conn,
		table: tableOrderItems,
	}
}

func (r *OrderItems) Create(ctx context.Context, order_items model.OrderItems) error {
	object := dao.FromOrderItems(order_items)
	query := "INSERT INTO " + r.table + " (order_id, product_id, quantity) VALUES ($1, $2, $3)"

	_, err := r.conn.Exec(query, object.OrderID, object.ProductID, object.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderItems) Get(ctx context.Context, id int) (model.OrderItems, error) {
	var order_items dao.OrderItems
	query := "SELECT order_id, product_id, quantity FROM " + r.table + " WHERE id = $1"

	err := r.conn.QueryRow(query, id).Scan(&order_items.OrderID, &order_items.ProductID, &order_items.Quantity)
	if err != nil {
		return model.OrderItems{}, err
	}

	return dao.ToOrderItems(order_items), nil
}

func (r *OrderItems) UPDATE(ctx context.Context, id int, order_items model.OrderItems) error {
	object := dao.FromOrderItems(order_items)
	query := "UPDATE " + r.table + "SET product_id = $1, quantity = $2 WHERE id = $3"

	_, err := r.conn.Exec(query, object.ProductID, object.Quantity, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderItems) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM " + r.table + " WHERE id = $1"

	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
