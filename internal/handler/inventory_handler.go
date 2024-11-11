package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/models"
	"hot-coffee/pkg/logger"
)

type InventoryHandler interface {
	AddInventoryItem(w http.ResponseWriter, r *http.Request)
	GetInventoryItems(w http.ResponseWriter, r *http.Request)
	GetInventoryItem(w http.ResponseWriter, r *http.Request)
	UpdateInventoryItem(w http.ResponseWriter, r *http.Request)
	DeleteInventoryItem(w http.ResponseWriter, r *http.Request)
	WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request)
}

type inventoryHandler struct {
	InventoryService service.InventoryService
	logger           *logger.Logger
}

func NewInventoryHandler(s service.InventoryService, l *logger.Logger) *inventoryHandler {
	return &inventoryHandler{InventoryService: s, logger: l}
}

func (h *inventoryHandler) WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
	}
}

func (h *inventoryHandler) WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	formattedJSON, err := json.MarshalIndent(jsonResponse, "", " ")
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	w.Write(formattedJSON)
}

func (h *inventoryHandler) WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
	// TODO: if its statusCode == 500 -> add ERROR log
	// TODO: in other cases 		  -> print DEBUG log
	// TODO: find case to add WARNING log (высосать из пальца)

	switch statusCode {
	case http.StatusInternalServerError:
		h.logger.PrintErrorMsg(err.Error())
	case http.StatusBadRequest,
		http.StatusNotFound,
		http.StatusUnsupportedMediaType,
		http.StatusConflict:

		h.logger.PrintDebugMsg(err.Error())
	}
	errorJSON := &models.ErrorResponse{Error: err.Error()}
	h.WriteJSONResponse(statusCode, errorJSON, w, r)
}

func (h *inventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Fix error of `{"error": "EOF"}`
	if r.Body == nil {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	var item models.InventoryItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		h.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	h.logger.PrintDebugMsg("Adding new inventory item: %+v", item)

	err := h.InventoryService.AddInventoryItem(item)
	if err != nil {
		switch err {
		case service.ErrNotUniqueID:
			h.WriteErrorResponse(http.StatusConflict, err, w, r)
			return
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.logger.PrintInfoMsg("Successfully added new inventory item: %+v", item)

	h.WriteJSONResponse(http.StatusCreated, item, w, r)
}

// 200 OK — запрос был успешно обработан.
// 201 Created — новый ресурс был успешно создан.
// 400 Bad Request — ошибка в запросе.
// 500 Internal Server Error — ошибка на сервере.

func (h *inventoryHandler) GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	data, err := h.InventoryService.RetrieveInventoryItems()
	if err != nil {
		switch err {
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.logger.PrintDebugMsg("Retrieved inventory items")

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *inventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	data, err := h.InventoryService.RetrieveInventoryItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.logger.PrintDebugMsg("Retrieved inventory item with ID: %s", itemId)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		h.logger.PrintErrorMsg("Failed to write response: %v", err)
	}
}

func (h *inventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	var item models.InventoryItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		h.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	err := h.InventoryService.UpdateInventoryItem(itemId, item)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		case service.ErrNotUniqueID:
			h.WriteErrorResponse(http.StatusBadRequest, fmt.Errorf("item with id '%s' not unique", itemId), w, r)
			return
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *inventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}
	err := h.InventoryService.DeleteInventoryItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.logger.PrintDebugMsg("Inventory item with ID: %s successfully deleted", itemId)
	w.WriteHeader(http.StatusNoContent)
}
