package service

import (
	"errors"
)

var (
	ErrNotValidIngredientID   error = errors.New("ingredient ID is not valid")
	ErrNotUniqueID            error = errors.New("ingredient ID must be unique")
	ErrNoItem                 error = errors.New("item not found")
	ErrNotValidIngredientName error = errors.New("ingredient name is not valid")
	ErrNotValidQuantity       error = errors.New("ingredient quantity is not valid")
	ErrNotValidUnit           error = errors.New("ingredient unit is not valid")

	ErrNotValidMenuID          error = errors.New("menu ID is not valid")
	ErrNotUniqueMenuID         error = errors.New("menu ID must be unique")
	ErrNotValidMenuName        error = errors.New("menu name is not valid")
	ErrNotValidMenuDescription error = errors.New("menu description cannot be empty")
	ErrNotValidPrice           error = errors.New("menu price must be greater than 0")

	ErrNotUniqueOrder error = errors.New("order ID must be unique")
)
