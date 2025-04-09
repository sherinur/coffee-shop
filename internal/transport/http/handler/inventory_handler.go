package handler

import (
	"errors"
	"god"
	"log/slog"
	"net/http"

	"coffee-shop/internal/model"
	dto "coffee-shop/internal/transport/dto/inventory"
	"coffee-shop/internal/transport/dto/response"
)

// type InventoryWriter interface {
// 	AddInventoryItem(*god.Context)
// 	UpdateInventoryItem(*god.Context)
// 	DeleteInventoryItem(*god.Context)
// }

// type InventoryReader interface {
// 	GetInventoryItems(*god.Context)
// 	GetInventoryItem(*god.Context)
// }

type InventoryHandler struct {
	service InventoryService
	log     *slog.Logger
}

func NewInventoryHandler(s InventoryService, l *slog.Logger) *InventoryHandler {
	return &InventoryHandler{service: s, log: l}
}

// AddInventoryItem handles the HTTP request to add a new inventory item.
// It processes the incoming request, validates the input, and interacts with the service layer to add the item.
// If successful, it returns the added item as a JSON response with a 201 status code.
func (h *InventoryHandler) AddInventoryItem(c *god.Context) {
	var item dto.InventoryRequest
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.service.AddInventoryItem(item.ToDomain())
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
func (h *InventoryHandler) GetInventoryItems(c *god.Context) {
	items, err := h.service.RetrieveInventoryItems()
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Retrieved inventory items")
	res := response.APIResponse{
		Status: http.StatusOK,
		Body:   god.H{"items": items},
	}
	c.JSON(res.Status, res)
}

// GetInventoryItem handles the HTTP request to retrieve a specific inventory item by its ID.
func (h *InventoryHandler) GetInventoryItem(c *god.Context) {
	id := c.Request.PathValue("id")
	item, err := h.service.RetrieveInventoryItem(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Retrieved inventory item with ID:", slog.String("itemId", id))
	res := response.APIResponse{
		Status: http.StatusOK,
		Body:   god.H{"item": item},
	}
	c.JSON(res.Status, res)
}

// UpdateInventoryItem handles the HTTP request to update an existing inventory item by its ID.
func (h *InventoryHandler) UpdateInventoryItem(c *god.Context) {
	itemId := c.Request.PathValue("id")

	var item model.Inventory
	err := c.ShouldBindJSON(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.service.UpdateInventoryItem(itemId, item)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Successfully updated an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusOK)
}

func (h *InventoryHandler) DeleteInventoryItem(c *god.Context) {
	itemId := c.Request.PathValue("id")

	err := h.service.DeleteInventoryItem(itemId)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.log.Debug("Successfully deleted an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusNoContent)
}

func (h *InventoryHandler) handleError(c *god.Context, err error) {
	var serviceErr *model.ServiceError
	if errors.As(err, &serviceErr) {
		c.JSON(serviceErr.Code, serviceErr.Hash())
		return
	}

	h.log.Error("Error of InventoryHandler", slog.String("error", err.Error()))
	c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
}
