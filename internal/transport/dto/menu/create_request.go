package dto

type MenuItemRequest struct {
	ID          string               `json:"product_id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       float64              `json:"price"`
	Ingredients []MenuItemIngredient `json:"ingredients"`
}

type MenuItemIngredient struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}

func (r *MenuItemRequest) Validate() error {
	// TODO: Write validation logic here
	if r.ID == "" {
		return nil
	}

	if r.Name == "" {
		return nil
	}

	if r.Description == "" {
		return nil
	}

	if r.Price < 0 {
		return nil
	}
	return nil
}
