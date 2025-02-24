package handler

import (
	"errors"
	"god"
	"log/slog"
	"net/http"

	"coffee-shop/internal/service"
	"coffee-shop/models"
)

type OrderWriter interface {
	CreateOrder(*god.Context)
	UpdateOrder(*god.Context)
	DeleteOrder(*god.Context)
	CloseOrder(*god.Context)
}

type OrderReader interface {
	RetrieveOrders(*god.Context)
	RetrieveOrder(*god.Context)
}

type orderHandler struct {
	OrderService service.OrderService
	log          *slog.Logger
}

func NewOrderHandler(s service.OrderService, l *slog.Logger) *orderHandler {
	return &orderHandler{OrderService: s, log: l}
}

func (h *orderHandler) CreateOrder(c *god.Context) {
	var order models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"code": http.StatusBadRequest, "error": err.Error(), "message": "Invalid request body"})
		return
	}

	h.log.Debug("Creating new order", slog.Any("Order", order))
	err = h.OrderService.AddOrder(order)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Info("Successfully created new order", slog.Any("Order", order))
	c.JSON(http.StatusCreated, god.H{"body": order})
}

func (h *orderHandler) RetrieveOrders(c *god.Context) {
	// Retrieve the orders from the service layer
	items, err := h.OrderService.RetrieveOrders()
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Retrieved orders")

	c.JSON(http.StatusOK, god.H{"body": items})
}

func (h *orderHandler) RetrieveOrder(c *god.Context) {
	id := c.Request.PathValue("id")
	item, err := h.OrderService.RetrieveOrder(id)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Retrieved order with ID", slog.String("OrderId", id))
	c.JSON(http.StatusOK, item)
}

func (h *orderHandler) UpdateOrder(c *god.Context) {
	id := c.Request.PathValue("id")
	var order models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.OrderService.UpdateOrder(id, order)
	if err != nil {
		h.handleError(c, err)
	}

	c.Status(http.StatusOK)
}

func (h *orderHandler) DeleteOrder(c *god.Context) {
	id := c.Request.PathValue("id")
	err := h.OrderService.DeleteOrder(id)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Successfully deleted order with ID ", slog.String("OrderId", id))
	c.Status(http.StatusNoContent)
}

func (h *orderHandler) CloseOrder(c *god.Context) {
	id := c.Request.PathValue("id")
	err := h.OrderService.CloseOrder(id)
	if err != nil {
		h.handleError(c, err)
	}
	c.Status(http.StatusOK)
}

func (h *orderHandler) handleError(c *god.Context, err error) {
	var serviceErr *service.ServiceError
	if errors.As(err, &serviceErr) {
		c.JSON(serviceErr.Code, serviceErr.Hash())
	} else {
		h.log.Error("Error of OrderHandler", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
	}
}
