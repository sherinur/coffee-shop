package service

import "errors"

var (
	ErrNotValidID          error = errors.New("ingredient id is not valid")
	ErrNotUniqueID         error = errors.New("ingredient ID must be unique")
	ErrNoItem              error = errors.New("item not found")
	ErrNotValidName        error = errors.New("ingredient name is not valid")
	ErrNotValidQuantity    error = errors.New("ingredient quantity is not valid")
	ErrNotValidUnit        error = errors.New("ingredient unit is not valid")
	ErrNotValidDescription error = errors.New("menu description is not valid")
	ErrNotValidPrice       error = errors.New("menu price is not valid")
	ErrNotValidIngredients error = errors.New("menu ingredients is not  valid")
)
