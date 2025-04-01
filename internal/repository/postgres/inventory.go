package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type Inventory struct {
	conn  *sql.DB
	table string
}

const (
	tableInventory = "inventory"
)

func NewInventory(conn *sql.DB) *Inventory {
	return &Inventory{
		conn:  conn,
		table: tableInventory,
	}
}

func (i *Inventory) Create(ctx context.Context, item model.Inventory) error {
	object := dao.FromInventory(item)
	query := "INSERT INTO " + i.table + " (name, quantity, unit) VALUES ($1, $2, $3)"

	_, err := i.conn.Exec(query, object.Name, object.Quantity, object.Unit)
	if err != nil {
		return err
	}

	return nil
}
