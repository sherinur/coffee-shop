package handler

import (
	"errors"
	"god"
	"log/slog"
	"net/http"

	"coffee-shop/internal/service"
	"coffee-shop/models"
)

type InventoryWriter interface {
	AddInventoryItem(*god.Context)
	UpdateInventoryItem(*god.Context)
	DeleteInventoryItem(*god.Context)
}

type InventoryReader interface {
	GetInventoryItems(*god.Context)
	GetInventoryItem(*god.Context)
}

type inventoryHandler struct {
	inventoryService service.InventoryService
	log              *slog.Logger
}

func NewInventoryHandler(s service.InventoryService, l *slog.Logger) *inventoryHandler {
	return &inventoryHandler{inventoryService: s, log: l}
}

// AddInventoryItem handles the HTTP request to add a new inventory item.
// It processes the incoming request, validates the input, and interacts with the service layer to add the item.
// If successful, it returns the added item as a JSON response with a 201 status code.
func (h *inventoryHandler) AddInventoryItem(c *god.Context) {
	var item models.InventoryItem
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.inventoryService.AddInventoryItem(item)
	h.log.Debug("Adding new inventory item:", slog.Any("item", item))
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Successfully added new inventory item:", slog.Any("item", item))
	c.JSON(http.StatusCreated, god.H{"item": item})
}

// GetInventoryItems handles the HTTP request to retrieve inventory items.
// It calls the service layer to get the list of inventory items, handles errors, and returns the data in the response.
func (h *inventoryHandler) GetInventoryItems(c *god.Context) {
	items, err := h.inventoryService.RetrieveInventoryItems()
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Retrieved inventory items")
	c.JSON(http.StatusOK, god.H{"items": items})
}

// GetInventoryItem handles the HTTP request to retrieve a specific inventory item by its ID.
func (h *inventoryHandler) GetInventoryItem(c *god.Context) {
	itemId := c.Request.PathValue("id")

	data, err := h.inventoryService.RetrieveInventoryItem(itemId)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Retrieved inventory item with ID:", slog.String("itemId", itemId))
	c.JSON(http.StatusOK, god.H{"item": data})
}

// UpdateInventoryItem handles the HTTP request to update an existing inventory item by its ID.
func (h *inventoryHandler) UpdateInventoryItem(c *god.Context) {
	itemId := c.Request.PathValue("id")

	var item models.InventoryItem
	err := c.ShouldBindJSON(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.inventoryService.UpdateInventoryItem(itemId, item)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Successfully updated an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusOK)
}

func (h *inventoryHandler) DeleteInventoryItem(c *god.Context) {
	itemId := c.Request.PathValue("id")

	err := h.inventoryService.DeleteInventoryItem(itemId)
	if err != nil {
		h.handleError(c, err)
	}

	h.log.Debug("Successfully deleted an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusNoContent)
}

func (h *inventoryHandler) handleError(c *god.Context, err error) {
	var serviceErr *service.ServiceError
	if errors.As(err, &serviceErr) {
		c.JSON(serviceErr.Code, serviceErr.Hash())
	} else {
		h.log.Error("Error of InventoryHandler", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
	}
}
