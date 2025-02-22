package handler

import (
	"fmt"
	"god"
	"log/slog"
	"net/http"

	"coffee-shop/internal/service"
	"coffee-shop/models"
)

type InventoryHandler interface {
	AddInventoryItem(*god.Context)
	GetInventoryItems(*god.Context)
	GetInventoryItem(*god.Context)
	UpdateInventoryItem(*god.Context)
	DeleteInventoryItem(*god.Context)
}

type inventoryHandler struct {
	InventoryService service.InventoryService
	log              *slog.Logger
}

func NewInventoryHandler(s service.InventoryService, l *slog.Logger) *inventoryHandler {
	return &inventoryHandler{InventoryService: s, log: l}
}

// AddInventoryItem handles the HTTP request to add a new inventory item.
// It processes the incoming request, validates the input, and interacts with the service layer to add the item.
// If successful, it returns the added item as a JSON response with a 201 status code.
func (h *inventoryHandler) AddInventoryItem(c *god.Context) {
	defer c.Request.Body.Close()
	var item models.InventoryItem
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": "invalid request body", "message": err.Error()})
		return
	}

	err = h.InventoryService.AddInventoryItem(item)
	h.log.Debug("Adding new inventory item:", slog.Any("item", item))
	if err != nil {
		switch err {
		case service.ErrNotUniqueID:
			c.JSON(http.StatusConflict, god.H{"error": err.Error(), "message": "item with the same ID already exists"})
			return
		case service.ErrNotValidIngredientID, service.ErrNotValidIngredientName, service.ErrNotValidQuantity, service.ErrNotValidUnit:
			c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
			return
		default:
			h.log.Error("Error of AddInventoryItem", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
			return
		}
	}

	h.log.Debug("Successfully added new inventory item:", slog.Any("item", item))
	c.JSON(http.StatusCreated, god.H{"item": item})
}

// GetInventoryItems handles the HTTP request to retrieve inventory items.
// It calls the service layer to get the list of inventory items, handles errors, and returns the data in the response.
func (h *inventoryHandler) GetInventoryItems(c *god.Context) {
	defer c.Request.Body.Close()
	items, err := h.InventoryService.RetrieveInventoryItems()
	if err != nil {
		switch err {
		default:
			h.log.Error("Error of GetInventoryItems", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
			return
		}
	}
	h.log.Debug("Retrieved inventory items")
	c.JSON(http.StatusOK, god.H{"items": items})
}

// GetInventoryItem handles the HTTP request to retrieve a specific inventory item by its ID.
func (h *inventoryHandler) GetInventoryItem(c *god.Context) {
	defer c.Request.Body.Close()
	itemId := c.Request.PathValue("id")

	data, err := h.InventoryService.RetrieveInventoryItem(itemId)
	if err != nil {
		switch err.Error() {
		case "identificator is not valid":
			c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request"})
		case "item not found":
			c.JSON(http.StatusNotFound, god.H{"error": err.Error(), "message": fmt.Sprintf("item with id '%s' not found", itemId)})
		default:
			h.log.Error("Error of GetInventoryItem", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
			return
		}
		return
	}

	h.log.Debug("Retrieved inventory item with ID:", slog.String("itemId", itemId))
	c.JSON(http.StatusOK, god.H{"item": data})
}

// UpdateInventoryItem handles the HTTP request to update an existing inventory item by its ID.
func (h *inventoryHandler) UpdateInventoryItem(c *god.Context) {
	defer c.Request.Body.Close()
	itemId := c.Request.PathValue("id")

	var item models.InventoryItem
	err := c.ShouldBindJSON(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		return
	}

	err = h.InventoryService.UpdateInventoryItem(itemId, item)
	if err != nil {
		if err.Error() == "identificator is not valid" {
			c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		}
		switch err {
		case service.ErrNoItem:
			c.JSON(http.StatusNotFound, god.H{"error": err.Error(), "message": fmt.Sprintf("item with id '%s' not found", itemId)})
			return
		case service.ErrNotUniqueID,
			service.ErrNotValidIngredientID,
			service.ErrNotValidIngredientName,
			service.ErrNotValidQuantity,
			service.ErrNotValidUnit:
			c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
			return
		default:
			h.log.Error("Error of UpdateInventoryItem", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
			return
		}
	}

	h.log.Debug("Successfully updated an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusOK)
}

func (h *inventoryHandler) DeleteInventoryItem(c *god.Context) {
	defer c.Request.Body.Close()
	itemId := c.Request.PathValue("id")

	err := h.InventoryService.DeleteInventoryItem(itemId)
	if err != nil {
		if err.Error() == "identificator is not valid" {
			c.JSON(http.StatusBadRequest, god.H{"error": err.Error(), "message": "invalid request body"})
		}
		switch err {
		case service.ErrNoItem:
			c.JSON(http.StatusNotFound, god.H{"error": err.Error(), "message": fmt.Sprintf("item with id '%s' not found", itemId)})
			return
		default:
			h.log.Error("Error of UpdateInventoryItem", slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, god.H{"error": err.Error(), "message": "internal server error"})
			return
		}
	}

	h.log.Debug("Successfully deleted an inventory item with ID:", slog.String("itemId", itemId))
	c.Status(http.StatusNoContent)
}
