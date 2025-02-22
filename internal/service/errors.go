package service

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

func NewServiceError(err error, code int, message string) *ServiceError {
	return &ServiceError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

var (
	// Inventory errors

	ErrNotValidIngredientID   error = NewServiceError(errors.New("invalid ingredient ID"), http.StatusBadRequest, "ingredient ID is not valid")
	ErrNotUniqueID            error = NewServiceError(errors.New("not unique ingredient ID"), http.StatusConflict, "item with the same ID already exists")
	ErrNoItem                 error = NewServiceError(errors.New("item not found"), http.StatusNotFound, "item with the given ID does not exist")
	ErrNotValidIngredientName error = NewServiceError(errors.New("invalid ingredient Name"), http.StatusBadRequest, "ingredient name is not valid")
	ErrNotValidQuantity       error = NewServiceError(errors.New("invalid ingredient Quantity"), http.StatusBadRequest, "ingredient quantity is not valid")
	ErrNotValidUnit           error = NewServiceError(errors.New("invalid ingredient Unit"), http.StatusBadRequest, "ingredient unit is not valid")

	// Menu errors

	ErrNotValidMenuID           error = errors.New("product ID is not valid")
	ErrNotUniqueMenuID          error = errors.New("product ID must be unique")
	ErrNotValidMenuName         error = errors.New("product name is not valid")
	ErrNotValidMenuDescription  error = errors.New("product description cannot be empty")
	ErrNotValidPrice            error = errors.New("product price must be greater than 0")
	ErrDuplicateMenuIngredients error = errors.New("the ingredients of the product must not be repeated")
	ErrNotValidIngredints       error = errors.New("product ingredients is not valid")

	// Order errors

	ErrNotValidOrderID           error = errors.New("order ID is not valid")
	ErrNotValidOrderCustomerName error = errors.New("order CustomeName is not valid")
	ErrDuplicateOrderItems       error = errors.New("the items in the order must not be repeated")
	ErrNotValidOrderItems        error = errors.New("order items is not valid ")
	ErrNotValidOrderProductID    error = errors.New("product ID is not valid")
	ErrNotValidStatusField       error = errors.New("status field cannot be set manually")
	ErrNotValidCreatedAt         error = errors.New("created_at field cannot be set manually")

	ErrNoOrder                    error = errors.New("order not found")
	ErrOrderProductNotFound       error = errors.New("product not found")
	ErrNotEnoughInventoryQuantity error = errors.New("not enough ingredient quantity")
	ErrProductNotFound            error = errors.New("the product is not on the menu")
	ErrInventoryItemNotFound      error = errors.New("ingredient not found")
	ErrOrderClosed                error = errors.New("order is closed")

	ErrNotUniqueOrder error = errors.New("order ID must be unique")
)
