package dao

import "coffee-shop/internal/model"

type MenuItem struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
}

func FromMenu(m model.MenuItem) MenuItem {
	return MenuItem{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
	}
}

func ToMenu(m MenuItem) model.MenuItem {
	return model.MenuItem{
		ID:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
	}
}

type MenuItemIngredients struct {
	MenuID       int `json:"menu_id" db:"menu_id"`
	IngredientID int `json:"ingredient_id" db:"ingredient_id"`
	Quantity     int `json:"quantity" db:"quantity"`
}

func FromIngredients(m model.MenuItemIngredients) MenuItemIngredients {
	return MenuItemIngredients{
		MenuID:       m.MenuID,
		IngredientID: m.IngredientID,
		Quantity:     m.Quantity,
	}
}

func ToIngredients(m MenuItemIngredients) model.MenuItemIngredients {
	return model.MenuItemIngredients{
		MenuID:       m.MenuID,
		IngredientID: m.IngredientID,
		Quantity:     m.Quantity,
	}
}
