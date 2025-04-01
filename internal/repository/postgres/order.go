package postgres

import "database/sql"

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
