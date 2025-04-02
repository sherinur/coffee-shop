package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type OrderStatusHistory struct {
	conn  *sql.DB
	table string
}

const (
	tableOrderStatusHistory = "order_status_history"
)

func NewOrderStatusHistory(conn *sql.DB) *OrderStatusHistory {
	return &OrderStatusHistory{
		conn:  conn,
		table: tableOrderStatusHistory,
	}
}

func (r *OrderStatusHistory) Create(ctx context.Context, order_history model.OrderStatusHistory) error {
	object := dao.FromOrderStatusHistory(order_history)
	query := "INSETR INTO " + r.table + "(order_id) VALUES ($1)"

	_, err := r.conn.Exec(query, object.OrderID)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderStatusHistory) Get(ctx context.Context, id int) (model.OrderStatusHistory, error) {
	var order_history dao.OrderStatusHistory
	query := "SELET id, order_id, opened_at, closed_at FROM " + r.table + " WHERE id = $1"

	err := r.conn.QueryRow(query, id).Scan(&order_history.ID, &order_history.OrderID, &order_history.OpenedAt, &order_history.ClosedAt)
	if err != nil {
		return model.OrderStatusHistory{}, err
	}

	return dao.ToOrderStatusHistory(order_history), nil
}

func (r *OrderStatusHistory) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM " + r.table + " WHERE id = $1"

	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
