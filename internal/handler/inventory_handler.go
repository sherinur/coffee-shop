package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"hot-coffee/internal/service"
	"hot-coffee/models"
	"log/slog"
	"net/http"
)

type InventoryHandler interface {
	AddInventoryItem(w http.ResponseWriter, r *http.Request)
	GetInventoryItems(w http.ResponseWriter, r *http.Request)
	GetInventoryItem(w http.ResponseWriter, r *http.Request)
	UpdateInventoryItem(w http.ResponseWriter, r *http.Request)
	DeleteInventoryItem(w http.ResponseWriter, r *http.Request)
}

type inventoryHandler struct {
	InventoryService service.InventoryService
}

func NewInventoryHandler(s service.InventoryService) *inventoryHandler {
	return &inventoryHandler{InventoryService: s}
}

func WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, err, w, r)
	}
}

func WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	formattedJSON, err := json.MarshalIndent(jsonResponse, "", " ")
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	w.Write(formattedJSON)
}

func WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
	// TODO: if its statusCode == 500 -> add ERROR log
	// TODO: in other cases 		  -> print DEBUG log
	// TODO: find case to add WARNING log (высосать из пальца)

	switch statusCode {
	case http.StatusInternalServerError:
		slog.Error(err.Error()) // example
	}

	errorJSON := &models.ErrorResponse{Error: err.Error()}
	WriteJSONResponse(statusCode, errorJSON, w, r)
}

func (h *inventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	defer r.Body.Close()

	var item models.InventoryItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	// TODO: Write debug log here

	err := h.InventoryService.AddInventoryItem(item)
	if err != nil {
		switch err {
		case service.ErrNotUniqueID:
			WriteErrorResponse(http.StatusConflict, err, w, r)
			return
		default:
			WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	WriteJSONResponse(http.StatusCreated, item, w, r)
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
			WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// TODO: Write debug log here
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *inventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	data, err := h.InventoryService.RetrieveInventoryItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// TODO: Write debug log here

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		// TODO: Print error log here
	}
}

func (h *inventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to update the inventory item by id.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEM updating by id"))
}

func (h *inventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
	}

	err := h.InventoryService.DeleteInventoryItem(itemId)
	if err != nil {
		switch err {
		case service.ErrNoItem:
			WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id %s not found", itemId), w, r)
			return
		default:
			WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// TODO: Write debug log here
	w.WriteHeader(http.StatusNoContent)
}
