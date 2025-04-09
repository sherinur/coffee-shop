package dto

import "coffee-shop/internal/model"

type InventoryRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

func (r *InventoryRequest) ToDomain() model.Inventory {
	return model.Inventory{
		Name:     r.Name,
		Quantity: r.Quantity,
		Unit:     r.Unit,
	}
}

func (r *InventoryRequest) Validate() error {
	if r.Name == "" {
		return nil
	}

	if r.Quantity < 0 {
		return nil
	}

	if r.Unit == "" {

	}
	return nil
}
