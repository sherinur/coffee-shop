package dto

import "coffee-shop/internal/model"

type InventoryResponse struct {
	IngredientID int    `json:"ingredient_id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	Unit         string `json:"unit"`
}

func NewInventoryResponse(i model.Inventory) InventoryResponse {
	return InventoryResponse{
		IngredientID: i.IngredientID,
		Name:         i.Name,
		Quantity:     i.Quantity,
		Unit:         i.Unit,
	}
}
