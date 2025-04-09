package model

import (
	"errors"
)

var (
	// Inventory errors

	ErrNotValidIngredientID   error = errors.New("invalid ingredient ID")
	ErrNotUniqueID            error = errors.New("not unique ingredient ID")
	ErrNoItem                 error = errors.New("item not found")
	ErrNotValidIngredientName error = errors.New("invalid ingredient Name")
	ErrNotValidQuantity       error = errors.New("invalid ingredient Quantity")
	ErrNotValidUnit           error = errors.New("invalid ingredient Unit")

	// Menu errors

	ErrNotValidMenuID           error = errors.New("invalid product ID")
	ErrNotUniqueMenuID          error = errors.New("not unique product ID")
	ErrNotValidMenuName         error = errors.New("invalid product Name")
	ErrNotValidMenuDescription  error = errors.New("invalid product Description")
	ErrNotValidPrice            error = errors.New("invalid product Price")
	ErrDuplicateMenuIngredients error = errors.New("invalid product Ingredients")
	ErrNotEnoughIngredients     error = errors.New("invalid product Ingredients")

	// Order errors

	ErrNotValidOrderID           error = errors.New("invalid order ID")
	ErrNotValidOrderCustomerName error = errors.New("invalid order CustomeName")
	ErrNotValidOrderStatus       error = errors.New("invalid order status")
	ErrNotValidOrderNotes        error = errors.New("invalid order notes")
	ErrDuplicateOrderItems       error = errors.New("duplicate order Items")
	ErrNotValidOrderItems        error = errors.New("invalid order Items")
	ErrNotValidOrderProductID    error = errors.New("invalid order product ID")
	ErrNotValidStatusField       error = errors.New("invalid order product ID")
	ErrNotValidCreatedAt         error = errors.New("invalid request")

	// Order status history errors

	ErrNotValidStatusHistoryTime error = errors.New("invalid time for history status")

	// Price history errors

	ErrNotValidPriceHistoryID error = errors.New("invalid price history id")
	ErrNotValidChangedAtTime  error = errors.New("invalid price history time")

	ErrNoOrder                    error = errors.New("order not found")
	ErrOrderProductNotFound       error = errors.New("product not found")
	ErrNotEnoughInventoryQuantity error = errors.New("invalid ingredient quantity")
	ErrProductNotFound            error = errors.New("product not found")
	ErrInventoryItemNotFound      error = errors.New("order not found")
	ErrOrderClosed                error = errors.New("order is closed")
	ErrNotUniqueOrder             error = errors.New("not unique order ID")
)
