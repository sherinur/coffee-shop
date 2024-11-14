package service

import (
	"errors"
)

var (
	ErrNoOrder error = errors.New("order not found")

	ErrNotValidIngredientID   error = errors.New("ingredient ID is not valid")
	ErrNotUniqueID            error = errors.New("ingredient ID must be unique")
	ErrNoItem                 error = errors.New("item not found")
	ErrNotValidIngredientName error = errors.New("ingredient name is not valid")
	ErrNotValidQuantity       error = errors.New("quantity is not valid")
	ErrNotValidUnit           error = errors.New("ingredient unit is not valid")

	ErrNotValidMenuID           error = errors.New("menu ID is not valid")
	ErrNotUniqueMenuID          error = errors.New("menu ID must be unique")
	ErrNotValidMenuName         error = errors.New("menu name is not valid")
	ErrNotValidMenuDescription  error = errors.New("menu description cannot be empty")
	ErrNotValidPrice            error = errors.New("menu price must be greater than 0")
	ErrDuplicateMenuIngredients error = errors.New("the ingredients in the menu should not be repeated")
	ErrNotValidIngredints       error = errors.New("menu ingredients is not valid ")

	ErrNotValidOrderID           error = errors.New("order ID is not valid")
	ErrNotValidOrderCustomerName error = errors.New("ordre CustomeName is not valid")
	ErrDuplicateOrderItems       error = errors.New("the items in the order should not be repeated")
	ErrNotValidOrderItems        error = errors.New("order items is not valid ")
	ErrNotValidOrderProductID    error = errors.New("product ID is not valid")
	ErrNotValidStatusField       error = errors.New("status cannot be set manually")
	ErrNotValidCreatedAt         error = errors.New("created_at cannot be set manually")

	ErrOrderProductNotFound       error = errors.New("product not found")
	ErrNotEnoughInventoryQuantity error = errors.New("not enough ingredient quantity")
	ErrProductNotFound            error = errors.New("the product is not on the menu")
	ErrInventoryItemNotFound      error = errors.New("ingredient not found")

	ErrNotUniqueOrder error = errors.New("order ID must be unique")
)
