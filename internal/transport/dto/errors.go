package dto

import "errors"

var (
	ErrNotValidMenuItemId      error = errors.New("menu item id cannot be empty")
	ErrNotValidMenuItemName    error = errors.New("menu name cannot be empty")
	ErrNotValidMenuDescription error = errors.New("menu description cannot be empty")
	ErrNotValidMenuPrice       error = errors.New("menu proce cannot be less than zero")
	ErrNotValidMenuIngredients error = errors.New("menu ingredients cannot be empty")

	ErrNotValdidMenuIngredientsQuantity error = errors.New("menu ingredient quantiry cannot be less or equal to zero")
)
