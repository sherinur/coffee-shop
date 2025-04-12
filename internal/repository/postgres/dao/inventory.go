package dao

import "coffee-shop/internal/model"

type Inventory struct {
	Id       int    `json:"ingredient_id" db:"ingredientid"`
	Name     string `json:"name" db:"name"`
	Quantity int    `json:"quantity" db:"quantity"`
	Unit     string `json:"unit" db:"unit"`
}

func FromInventory(item model.Inventory) Inventory {
	return Inventory{
		Name:     item.Name,
		Quantity: item.Quantity,
		Unit:     item.Unit,
	}
}

func ToInventory(item Inventory) model.Inventory {
	return model.Inventory{
		IngredientID: item.Id,
		Name:         item.Name,
		Quantity:     item.Quantity,
		Unit:         item.Unit,
	}
}
