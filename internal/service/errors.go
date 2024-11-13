package service

import (
	"errors"
)

var (
	ErrNotValidID       error = errors.New("ingredient id is not valid")
	ErrNotUniqueID      error = errors.New("ingredient ID must be unique")
	ErrNoItem           error = errors.New("item not found")
	ErrNotValidName     error = errors.New("ingredient name is not valid")
	ErrNotValidQuantity error = errors.New("ingredient quantity is not valid")
	ErrNotValidUnit     error = errors.New("ingredient unit is not valid")

	ErrNotValidDescription  error = errors.New("menu description is not valid")
	ErrNotValidPrice        error = errors.New("menu price must be greater than 0")
	ErrNotValidIngredients  error = errors.New("menu ingredients must be greater than 1")
	ErrIDContainsSpace      error = errors.New("ID cannot contain spaces")
	ErrInvalidIngredientID  error = errors.New("ingredient ID cannot be empty or contain spaces")
	ErrInvalidIngredientQty error = errors.New("ingredient quantity must be greater than 1")

	ErrNotUniqueOrder error = errors.New("order ID must be unique")
)
