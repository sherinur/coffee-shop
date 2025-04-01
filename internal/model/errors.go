package model

import (
	"errors"
	"fmt"
	"net/http"
)

// ServiceError is a wrapper struct of error for the service level
type ServiceError struct {
	Err     error
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("service error: %v (status: %d)", e.Err, e.Code)
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

// Hash returns the hash map of the ServiceError
func (e *ServiceError) Hash() map[string]any {
	return map[string]any{
		"code":    e.Code,
		"error":   e.Err.Error(),
		"message": e.Message,
	}
}

func NewServiceError(errStr string, code int, message string) *ServiceError {
	return &ServiceError{
		Err:     errors.New(errStr),
		Code:    code,
		Message: message,
	}
}

var (
	// Inventory errors

	ErrNotValidIngredientID   error = NewServiceError("invalid ingredient ID", http.StatusBadRequest, "ingredient ID is not valid")
	ErrNotUniqueID            error = NewServiceError("not unique ingredient ID", http.StatusConflict, "item with the same ID already exists")
	ErrNoItem                 error = NewServiceError("item not found", http.StatusNotFound, "item with the given ID does not exist")
	ErrNotValidIngredientName error = NewServiceError("invalid ingredient Name", http.StatusBadRequest, "ingredient name is not valid")
	ErrNotValidQuantity       error = NewServiceError("invalid ingredient Quantity", http.StatusBadRequest, "ingredient quantity is not valid")
	ErrNotValidUnit           error = NewServiceError("invalid ingredient Unit", http.StatusBadRequest, "ingredient unit is not valid")

	// Menu errors

	ErrNotValidMenuID           error = NewServiceError("invalid product ID", http.StatusBadRequest, "product ID is not valid")
	ErrNotUniqueMenuID          error = NewServiceError("not unique product ID", http.StatusConflict, "product with the same ID already exists")
	ErrNotValidMenuName         error = NewServiceError("invalid product Name", http.StatusBadRequest, "product name is not valid")
	ErrNotValidMenuDescription  error = NewServiceError("invalid product Description", http.StatusBadRequest, "product description cannot be empty")
	ErrNotValidPrice            error = NewServiceError("invalid product Price", http.StatusBadRequest, "product price must be greater than 0")
	ErrDuplicateMenuIngredients error = NewServiceError("invalid product Ingredients", http.StatusBadRequest, "ingredients of the product must not be repeated")
	ErrNotEnoughIngredients     error = NewServiceError("invalid product Ingredients", http.StatusBadRequest, "product must contain at least 1 ingredient")

	// Order errors

	ErrNotValidOrderID           error = NewServiceError("invalid order ID", http.StatusBadRequest, "order ID is not valid")
	ErrNotValidOrderCustomerName error = NewServiceError("invalid order CustomeName", http.StatusBadRequest, "order CustomeName is not valid")
	ErrNotValidOrderStatus       error = NewServiceError("invalid order status", http.StatusBadRequest, "order status is not valid")
	ErrNotValidOrderNotes        error = NewServiceError("invalid order notes", http.StatusBadRequest, "orders notes is not valid")
	ErrDuplicateOrderItems       error = NewServiceError("duplicate order Items", http.StatusBadRequest, "the items in the order must not be repeated")
	ErrNotValidOrderItems        error = NewServiceError("invalid order Items", http.StatusBadRequest, "order items is not valid")
	ErrNotValidOrderProductID    error = NewServiceError("invalid order product ID", http.StatusBadRequest, "product ID is not valid")
	ErrNotValidStatusField       error = NewServiceError("invalid order product ID", http.StatusBadRequest, "product ID is not valid")
	ErrNotValidCreatedAt         error = NewServiceError("invalid request", http.StatusBadRequest, "created_at field cannot be set manually")

	// Order status history errors

	ErrNotValidStatusHistoryTime error = NewServiceError("invalid time for history status", http.StatusBadRequest, "invalid time for history status")

	// Price history errors

	ErrNotValidPriceHistoryID error = NewServiceError("invalid price history id", http.StatusBadRequest, "price history id is not valid")
	ErrNotValidChangedAtTime  error = NewServiceError("invalid price history time", http.StatusBadRequest, "price history time is not valid")

	ErrNoOrder                    error = NewServiceError("order not found", http.StatusNotFound, "order not found")
	ErrOrderProductNotFound       error = NewServiceError("product not found", http.StatusNotFound, "product not found")
	ErrNotEnoughInventoryQuantity error = NewServiceError("invalid ingredient quantity", http.StatusBadRequest, "not enough ingredient quantity")
	ErrProductNotFound            error = NewServiceError("product not found", http.StatusBadRequest, "the product is not on the menu")
	ErrInventoryItemNotFound      error = NewServiceError("order not found", http.StatusBadRequest, "ingredient not found")
	ErrOrderClosed                error = NewServiceError("order is closed", http.StatusBadRequest, "can not edit the closed order")
	ErrNotUniqueOrder             error = NewServiceError("not unique order ID", http.StatusConflict, "order ID must be unique")
)
