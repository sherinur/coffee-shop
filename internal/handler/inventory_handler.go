package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/models"
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

func NewInventoryHandler(s service.InventoryService) InventoryHandler {
	return &inventoryHandler{InventoryService: s}
}

// 200 OK — запрос был успешно обработан.
// 201 Created — новый ресурс был успешно создан.
// 400 Bad Request — ошибка в запросе.
// 500 Internal Server Error — ошибка на сервере.

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

	if statusCode == 500 {
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

	var item models.InventoryItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	// TODO: Write debug log here

	if err := h.InventoryService.AddInventoryItem(item); err != nil {
		WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	WriteJSONResponse(http.StatusCreated, item, w, r)
}

func (h *inventoryHandler) GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to RETRIEVE ALL ITEMS.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEMS RETRIEVING"))
}

func (h *inventoryHandler) GetInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to get the inventory item by id.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEM retrieving by id"))
}

func (h *inventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to update the inventory item by id.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEM updating by id"))
}

func (h *inventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to delete the inventory item by id.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEM deleting by id"))
}
