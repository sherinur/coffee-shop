package postgres

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository/postgres/dao"
	"context"
	"database/sql"
)

type MenuItemIngredients struct {
	conn  *sql.DB
	table string
}

const (
	tableMenuItemIngredients = "menu_item_ingredients"
)

func NewMenuItemIngredients(conn *sql.DB) *MenuItemIngredients {
	return &MenuItemIngredients{
		conn:  conn,
		table: tableMenuItemIngredients,
	}
}

func (r *MenuItemIngredients) Create(ctx context.Context, menu_ingredients model.MenuItemIngredients) error {
	object := dao.FromIngredients(menu_ingredients)
	query := "INSERT INTO " + r.table + " (menu_id, ingredient_id, quantity) VALUES $1, $2, $3"

	_, err := r.conn.Exec(query, object.MenuID, object.IngredientID, object.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *MenuItemIngredients) Get(ctx context.Context, id int) (model.MenuItemIngredients, error) {
	var menu_ingredients dao.MenuItemIngredients
	query := "SELECT menu_id, ingredient_id, quantity FROM " + r.table + "WHERE id = $1"

	err := r.conn.QueryRow(query, id).Scan(&menu_ingredients.MenuID, &menu_ingredients.IngredientID, &menu_ingredients.Quantity)
	if err != nil {
		return model.MenuItemIngredients{}, nil
	}

	return dao.ToIngredients(menu_ingredients), nil
}

func (r *MenuItemIngredients) UPDATE(ctx context.Context, id int, menu_ingredients model.MenuItemIngredients) error {
	object := dao.FromIngredients(menu_ingredients)
	query := "UPDATE " + r.table + " SET menu_id = $1, , ingredient_id = $2, quantity = $3 WHERE id = $4"

	_, err := r.conn.Exec(query, object.MenuID, object.IngredientID, object.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (r *MenuItemIngredients) DELETE(ctx context.Context, id int) error {
	query := "DELETE FROM " + r.table + " WHERE id = $1"

	_, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
