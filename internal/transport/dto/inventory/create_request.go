package dto

type InventoryItemRequest struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

func (r *InventoryItemRequest) Validate() error {
	// TODO: Write validation logic here
	if r.IngredientID == "" {
		return nil
	}

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
