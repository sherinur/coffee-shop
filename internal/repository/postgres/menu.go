package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type Menu struct {
	conn  *sql.DB
	table string
}

const (
	tableMenu = "menu"
)

func NewMenu(conn *sql.DB) *Menu {
	return &Menu{
		conn:  conn,
		table: tableMenu,
	}
}

func (r *Menu) Create(ctx context.Context, menu model.MenuItem) error {
	object := dao.FromMenu(menu)
	query := "INSERT INTO " + r.table + " (name, description, price) VALUES ($1, $2, $3)"

	_, err := r.conn.Exec(query, object.Name, object.Description, object.Price)
	if err != nil {
		return nil
	}

	return nil
}

func (r *Menu) Get(ctx context.Context, id int) (model.MenuItem, error) {
	var menu dao.MenuItem
	query := "SELECT id, name, description, price FROM " + r.table + " WHERE id = $1"

	err := r.conn.QueryRow(query, id).Scan(&menu.Id, &menu.Name, &menu.Description, &menu.Price)
	if err != nil {
		return model.MenuItem{}, nil
	}

	return dao.ToMenu(menu), nil
}

func (r *Menu) GetAll(ctx context.Context) ([]model.MenuItem, error) {
	var menu_all []model.MenuItem
	query := "SELECT * FROM " + r.table

	rows, err := r.conn.Query(query)
	if err != nil {
		return []model.MenuItem{}, err
	}

	for rows.Next() {
		var menu_item dao.MenuItem
		err := rows.Scan(&menu_item.Id, &menu_item.Name, &menu_item.Description, &menu_item.Price)
		if err != nil {
			return []model.MenuItem{}, nil
		}

		menu_all = append(menu_all, dao.ToMenu(menu_item))
	}

	return menu_all, nil
}

func (r *Menu) Update(ctx context.Context, id int, menu model.MenuItem) error {
	object := dao.FromMenu(menu)
	query := "UPDATE " + r.table + " SET name = $1, description = $2, price = $3 WHERE id = $4"

	_, err := r.conn.Exec(query, object.Name, object.Description, object.Price, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Menu) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM " + r.table + " WHERE id = $1"

	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
