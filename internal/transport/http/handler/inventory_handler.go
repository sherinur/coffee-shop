package handler

import (
	"context"
	"errors"
	"god"
	"log/slog"
	"net/http"
	"strconv"

	"coffee-shop/internal/service"
	dto "coffee-shop/internal/transport/dto/inventory"
	"coffee-shop/internal/transport/dto/response"
)

type InventoryHandler interface {
	AddInventoryItem(*god.Context)
	GetInventoryItems(*god.Context)
	GetInventoryItem(*god.Context)
	UpdateInventoryItem(*god.Context)
	DeleteInventoryItem(*god.Context)
}

type inventoryHandler struct {
	service InventoryService
	log     *slog.Logger
}

func NewInventoryHandler(s InventoryService, l *slog.Logger) *inventoryHandler {
	return &inventoryHandler{service: s, log: l}
}

// AddInventoryItem handles the HTTP request to add a new inventory item.
// It processes the incoming request, validates the input, and interacts with the service layer to add the item.
// If successful, it returns the added item as a JSON response with a 201 status code.
func (h *inventoryHandler) AddInventoryItem(c *god.Context) {
	var item dto.InventoryRequest
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = item.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error()})
	}

	err = h.service.AddInventoryItem(context.TODO(), item.ToDomain())
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Successfully added new inventory item:", slog.Any("item", item))
	res := response.APIResponse{
		Status: http.StatusCreated,
		Body:   god.H{"item": item},
	}
	c.JSON(res.Status, res)
}

// GetInventoryItems handles the HTTP request to retrieve inventory items.
// It calls the service layer to get the list of inventory items, handles errors, and returns the data in the response.
func (h *inventoryHandler) GetInventoryItems(c *god.Context) {
	object, err := h.service.RetrieveInventoryItems(context.TODO())
	if err != nil {
		h.handleError(c, err)
		return
	}

	var items []dto.InventoryResponse
	for _, i := range object {
		items = append(items, dto.NewInventoryResponse(i))
	}

	h.log.Debug("Retrieved inventory items")
	res := response.APIResponse{
		Status: http.StatusOK,
		Body:   god.H{"items": items},
	}
	c.JSON(res.Status, res)
}

// GetInventoryItem handles the HTTP request to retrieve a specific inventory item by its ID.
func (h *inventoryHandler) GetInventoryItem(c *god.Context) {
	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	object, err := h.service.RetrieveInventoryItem(context.TODO(), itemID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	item := dto.NewInventoryResponse(*object)

	h.log.Debug("Retrieved inventory item with ID:", slog.String("itemId", id))
	res := response.APIResponse{
		Status: http.StatusOK,
		Body:   god.H{"item": item},
	}
	c.JSON(res.Status, res)
}

// UpdateInventoryItem handles the HTTP request to update an existing inventory item by its ID.
func (h *inventoryHandler) UpdateInventoryItem(c *god.Context) {
	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	var item dto.InventoryRequest
	err = c.ShouldBindJSON(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.service.UpdateInventoryItem(context.TODO(), itemID, item.ToDomain())
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Successfully updated an inventory item with ID:", slog.String("itemId", id))
	c.Status(http.StatusOK)
}

func (h *inventoryHandler) DeleteInventoryItem(c *god.Context) {
	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	err = h.service.DeleteInventoryItem(context.TODO(), itemID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Successfully deleted an inventory item with ID:", slog.String("itemId", id))
	c.Status(http.StatusNoContent)
}

func (h *inventoryHandler) handleError(c *god.Context, err error, code int) {
	var serviceErr *service.ServiceError
	if errors.As(err, &serviceErr) {
		c.JSON(serviceErr.Code, serviceErr.Hash())
		return
	}

	h.log.Error("Error of InventoryHandler", slog.String("error", err.Error()))
	c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
}
