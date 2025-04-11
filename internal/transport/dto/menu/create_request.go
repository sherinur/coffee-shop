package dto

import (
	"coffee-shop/internal/model"
	"coffee-shop/internal/transport/dto"
)

type MenuItemRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       float64              `json:"price"`
	Ingredients []MenuItemIngredient `json:"ingredients"`
}

type MenuItemIngredient struct {
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}

func (r *MenuItemRequest) Validate() error {
	switch {
	case r.Name == "":
		return dto.ErrNotValidMenuItemName
	case r.Description == "":
		return dto.ErrNotValidMenuDescription
	case r.Price <= 0:
		return dto.ErrNotValidMenuPrice
	case len(r.Ingredients) == 0:
		return dto.ErrNotValidMenuIngredients
	}

	for _, i := range r.Ingredients {
		i.Validate()
	}

	return nil
}

func (r *MenuItemIngredient) Validate() error {
	switch {
	case r.Quantity <= 0:
		return dto.ErrNotValdidMenuIngredientsQuantity
	default:
		return nil
	}
}

func ToDomain(m MenuItemRequest) (model.MenuItem, []model.MenuItemIngredients) {
	menuItem := model.MenuItem{
		ID:          0,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
	}

	var ingredients []model.MenuItemIngredients
	for _, ingredient := range m.Ingredients {
		ingredients = append(ingredients, model.MenuItemIngredients{
			MenuID:       0,
			IngredientID: ingredient.IngredientID,
			Quantity:     ingredient.Quantity,
		})
	}

	return menuItem, ingredients
}
