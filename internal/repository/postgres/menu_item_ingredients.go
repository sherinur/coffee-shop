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

func (r *MenuItemIngredients) GetAllWithID(ctx context.Context, id int) ([]model.MenuItemIngredients, error) {
	var ingredients []model.MenuItemIngredients
	query := "SELECT menu_id, ingredient_id, quantity FROM " + r.table + "WHERE menu_id = $1"

	rows, err := r.conn.Query(query)
	if err != nil {
		return []model.MenuItemIngredients{}, err
	}

	for rows.Next() {
		var ing dao.MenuItemIngredients
		err := rows.Scan(&ing.MenuID, &ing.IngredientID, &ing.Quantity)
		if err != nil {
			return []model.MenuItemIngredients{}, err
		}

		ingredients = append(ingredients, dao.ToIngredients(ing))
	}

	return ingredients, nil
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
