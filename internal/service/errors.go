package service

import "errors"

var (
	ErrNotValidID       error = errors.New("ingredient id is not valid")
	ErrNotUniqueID      error = errors.New("ingredient ID must be unique")
	ErrNoItem           error = errors.New("item not found")
	ErrNotValidName     error = errors.New("ingredient name is not valid")
	ErrNotValidQuantity error = errors.New("ingredient quantity is not valid")
	ErrNotValidUnit     error = errors.New("ingredient unit is not valid")

	ErrNotUniqueOrder error = errors.New("order ID must be unique")
)
