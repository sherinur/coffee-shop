package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"coffee-shop/internal/service"
	"coffee-shop/internal/utils"
	"coffee-shop/models"
)

type OrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	RetrieveOrders(w http.ResponseWriter, r *http.Request)
	RetrieveOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	CloseOrder(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	OrderService service.OrderService
	log          *slog.Logger
}

func NewOrderHandler(s service.OrderService, l *slog.Logger) *orderHandler {
	return &orderHandler{OrderService: s, log: l}
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	var order models.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		if err == io.EOF {
			utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
			return
		}
		utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	h.log.Debug("Creating new order", slog.Any("Order", order))

	err := h.OrderService.AddOrder(order)
	if err != nil {
		switch err {
		case service.ErrNotUniqueOrder:
			utils.WriteErrorResponse(http.StatusConflict, err, w, r)
			return
		case service.ErrNotValidOrderID,
			service.ErrNotValidOrderCustomerName,
			service.ErrNotValidStatusField,
			service.ErrNotValidCreatedAt,
			service.ErrNotValidOrderItems,
			service.ErrNotValidIngredientID,
			service.ErrDuplicateOrderItems,
			service.ErrNotValidQuantity,
			service.ErrNotValidOrderProductID,
			service.ErrNotEnoughInventoryQuantity:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
			return
		case service.ErrOrderProductNotFound,
			service.ErrInventoryItemNotFound:
			utils.WriteErrorResponse(http.StatusUnprocessableEntity, err, w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Info("Successfully created new order", slog.Any("Order", order))

	utils.WriteJSONResponse(http.StatusCreated, order, w, r)
}

func (h *orderHandler) RetrieveOrders(w http.ResponseWriter, r *http.Request) {
	// Retrieve the orders from the service layer
	data, err := h.OrderService.RetrieveOrders()
	if err != nil {
		switch err {
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Debug("Retrieved orders")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *orderHandler) RetrieveOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	if len(orderId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("order id is not valid"), w, r)
		return
	}

	data, err := h.OrderService.RetrieveOrder(orderId)
	if err != nil {
		if err.Error() == "order not found" {
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("order with id '%s' not found", orderId), w, r)
			return
		} else {
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Debug("Retrieved order with ID", slog.String("OrderId", orderId))

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		h.log.Error("Failed to write response", "error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *orderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	orderId := r.PathValue("id")
	if len(orderId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	var order models.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	err := h.OrderService.UpdateOrder(orderId, order)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("order with id '%s' not found", orderId), w, r)
			return
		case service.ErrNotValidOrderID,
			service.ErrNotValidOrderCustomerName,
			service.ErrNotValidStatusField,
			service.ErrNotValidCreatedAt,
			service.ErrNotValidOrderItems,
			service.ErrNotValidIngredientID,
			service.ErrDuplicateOrderItems,
			service.ErrNotValidQuantity,
			service.ErrNotValidOrderProductID,
			service.ErrOrderProductNotFound,
			service.ErrNotEnoughInventoryQuantity,
			service.ErrInventoryItemNotFound:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *orderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	if len(orderId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("order id is not valid"), w, r)
		return
	}

	err := h.OrderService.DeleteOrder(orderId)
	if err != nil {
		if err.Error() == "order not found" {
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("order with id '%s' not found", orderId), w, r)
			return
		}
	}

	h.log.Debug("Successfully deleted order with ID ", slog.String("OrderId", orderId))

	w.WriteHeader(http.StatusNoContent)
}

func (h *orderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("id")

	if len(orderId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("order id is not valid"), w, r)
		return
	}

	err := h.OrderService.CloseOrder(orderId)
	if err != nil {
		switch err {
		default:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		}
	}

	w.WriteHeader(http.StatusOK)
}
