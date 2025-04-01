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

func (i *Inventory) Get(ctx context.Context, id int) (model.Inventory, error) {
	var item dao.Inventory
	query := "SELECT id, name, quantity, unit FROM " + i.table + " WHERE id = $1"

	err := i.conn.QueryRow(query, id).Scan(&item.Id, &item.Name, &item.Quantity, &item.Unit)
	if err != nil {
		return model.Inventory{}, err
	}

	return dao.ToInventory(item), nil
}

func (i *Inventory) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM " + i.table + " WHERE id = $1"

	_, err := i.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
