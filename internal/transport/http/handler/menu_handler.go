package handler

import (
	"context"
	"god"
	"log/slog"
	"net/http"
	"strconv"

	dto "coffee-shop/internal/transport/dto/menu"
)

type MenuItem interface {
	AddMenuItem(*god.Context)
	UpdateMenuItem(*god.Context)
	GetAllMenuItems(c *god.Context)
	GetMenuItem(*god.Context)
	DeleteMenuItem(*god.Context)
}

type menuHandler struct {
	service MenuService
	log     *slog.Logger
}

func NewMenuHandler(s MenuService, l *slog.Logger) *menuHandler {
	return &menuHandler{service: s, log: l}
}

// AddMenuItem handles the HTTP request to add a new menu item.
// It processes the request body, validates the input, and calls the service layer to add the item.
func (h *menuHandler) AddMenuItem(c *god.Context) {
	var menu dto.MenuItemRequest
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"code": http.StatusBadRequest, "error": err.Error(), "message": "Invalid request body"})
		return
	}
	h.log.Debug("Adding new menu item", slog.Any("MenuItem", menu))

	item, ingredients := dto.ToDomain(menu)
	err = h.service.AddMenuItem(context.TODO(), item, ingredients)
	if err != nil {
		h.handleError(c, err, 400)
		return
	}

	h.log.Info("Successfully added new menu item", slog.Any("MenuItem", item))
	c.JSON(http.StatusCreated, god.H{"code": http.StatusCreated, "message": "Menu item added successfully"})
}

// GetMenuItems handles the HTTP request to retrieve all menu items.
// It calls the service layer to fetch the data and returns it to the client.
func (h *menuHandler) GetAllMenuItems(c *god.Context) {
	items, err := h.service.RetrieveMenuItems(context.TODO())
	if err != nil {
		h.handleError(c, err, 400)
		return
	}
	h.log.Debug("Retrieved Menu items")
	c.JSON(http.StatusOK, god.H{"code": http.StatusOK, "body": items})
}

// GetMenuItem handles the HTTP request to retrieve a specific menu item by its ID.
// It checks if the item ID is valid, calls the service layer to fetch the menu item,
// and returns the result to the client. In case of errors, it responds with the appropriate error message.
func (h *menuHandler) GetMenuItem(c *god.Context) {
	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err, 400)
		return
	}

	item, ingredient, err := h.service.RetrieveMenuItemWithId(context.TODO(), itemID)
	if err != nil {
		h.handleError(c, err, 400)
		return
	}

	menu := dto.NewMenuItemResponse(item, ingredient)

	h.log.Debug("Retrieved menu item with ID", slog.String("id", id))
	c.JSON(http.StatusOK, god.H{"code": http.StatusOK, "body": menu})
}

// UpdateMenuItem handles the HTTP request to update an existing menu item.
// It checks if the request body is valid, decodes the new menu item data, and
// calls the service layer to update the menu item. In case of errors, it responds
// with the appropriate HTTP status and error message.
func (h *menuHandler) UpdateMenuItem(c *god.Context) {
	var menu dto.MenuItemRequest
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, god.H{"code": http.StatusBadRequest, "error": err.Error(), "message": "Invalid request body"})
		return
	}

	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err, 400)
		return
	}

	item, ingredients := dto.ToDomain(menu)
	err = h.service.UpdateMenuItem(context.TODO(), itemID, item, ingredients)
	if err != nil {
		h.handleError(c, err, 400)
	}

	c.Status(http.StatusOK)
}

// DeleteMenuItem handles the HTTP request to delete a menu item by its ID.
// It validates the item ID, calls the service layer to delete the item, and
// responds with the appropriate HTTP status and message.
func (h *menuHandler) DeleteMenuItem(c *god.Context) {
	id := c.Request.PathValue("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		h.handleError(c, err, 400)
	}

	err = h.service.DeleteMenuItem(context.TODO(), itemID)
	if err != nil {
		h.handleError(c, err, 400)
	}

	h.log.Debug("Successfully deleted a menu item with ID ", slog.String("id", id))
	c.Status(http.StatusNoContent)
}

func (h *menuHandler) handleError(c *god.Context, err error, code int) {
	c.JSON(code, god.H{"error": err.Error()})
}
