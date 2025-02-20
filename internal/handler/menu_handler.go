package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
)

type MenuHandler interface {
	AddMenuItem(w http.ResponseWriter, r *http.Request)
	GetMenuItems(w http.ResponseWriter, r *http.Request)
	GetMenuItem(w http.ResponseWriter, r *http.Request)
	UpdateMenuItem(w http.ResponseWriter, r *http.Request)
	DeleteMenuItem(w http.ResponseWriter, r *http.Request)
}

type menuHandler struct {
	MenuService service.MenuService
	log         *slog.Logger
}

func NewMenuHandler(s service.MenuService, l *slog.Logger) *menuHandler {
	return &menuHandler{MenuService: s, log: l}
}

// AddMenuItem handles the HTTP request to add a new menu item.
// It processes the request body, validates the input, and calls the service layer to add the item.
func (h *menuHandler) AddMenuItem(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	var item models.MenuItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		switch err {
		case io.EOF:
			utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		default:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		}
		return
	}

	h.log.Debug("Adding new menu item", slog.Any("MenuItem", item))

	err := h.MenuService.AddMenuItem(item)
	if err != nil {
		switch err {
		case service.ErrNotUniqueMenuID:
			utils.WriteErrorResponse(http.StatusConflict, err, w, r)
			return
		case service.ErrNotValidMenuID,
			service.ErrNotValidMenuName,
			service.ErrNotValidMenuDescription,
			service.ErrNotValidPrice,
			service.ErrNotValidIngredientID,
			service.ErrNotValidQuantity,
			service.ErrDuplicateMenuIngredients,
			service.ErrNotValidIngredints:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Info("Successfully added new menu item", slog.Any("MenuItem", item))

	utils.WriteJSONResponse(http.StatusCreated, item, w, r)
}

// GetMenuItems handles the HTTP request to retrieve all menu items.
// It calls the service layer to fetch the data and returns it to the client.
func (h *menuHandler) GetMenuItems(w http.ResponseWriter, r *http.Request) {
	data, err := h.MenuService.RetrieveMenuItems()
	if err != nil {
		switch err {
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Debug("Retrieved Menu items")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetMenuItem handles the HTTP request to retrieve a specific menu item by its ID.
// It checks if the item ID is valid, calls the service layer to fetch the menu item,
// and returns the result to the client. In case of errors, it responds with the appropriate error message.
func (h *menuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")

	if len(itemId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	data, err := h.MenuService.RetrieveMenuItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Debug("Retrieved menu item with ID", slog.String("ItemId", itemId))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		h.log.Error("Failed to write response", "error", err)
	}
}

// UpdateMenuItem handles the HTTP request to update an existing menu item.
// It checks if the request body is valid, decodes the new menu item data, and
// calls the service layer to update the menu item. In case of errors, it responds
// with the appropriate HTTP status and error message.
func (h *menuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	var item models.MenuItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	err := h.MenuService.UpdateMenuItem(itemId, item)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		case service.ErrNotUniqueMenuID,
			service.ErrNotValidMenuID,
			service.ErrNotValidMenuName,
			service.ErrNotValidMenuDescription,
			service.ErrNotValidPrice,
			service.ErrNotValidIngredientID,
			service.ErrNotValidQuantity,
			service.ErrDuplicateMenuIngredients:
			utils.WriteErrorResponse(http.StatusBadRequest, err, w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteMenuItem handles the HTTP request to delete a menu item by its ID.
// It validates the item ID, calls the service layer to delete the item, and
// responds with the appropriate HTTP status and message.
func (h *menuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		utils.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	err := h.MenuService.DeleteMenuItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			utils.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.log.Debug("Successfully deleted a menu item with ID ", slog.String("ItemId", itemId))

	w.WriteHeader(http.StatusNoContent)
}
