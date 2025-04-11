package dto

import "coffee-shop/internal/model"

type MenuItemResponse struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Ingredients []MenuItemIngredients `json:"ingredients"`
	Price       float64               `json:"price"`
}

type MenuItemIngredients struct {
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}

func NewMenuItemResponse(m *model.MenuItem, i []model.MenuItemIngredients) MenuItemResponse {
	var ingredients []MenuItemIngredients
	for _, ing := range i {
		ingredients = append(ingredients, MenuItemIngredients{
			IngredientID: ing.IngredientID,
			Quantity:     ing.Quantity,
		})
	}

	menu := MenuItemResponse{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Ingredients: ingredients,
		Price:       m.Price,
	}
	return menu
}
