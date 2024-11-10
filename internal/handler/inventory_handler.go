package handler

import (
	"encoding/json"
	"fmt"
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

func WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}

func WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
	// TODO: Add ERROR log here, when error is writing
	errorJSON := &models.ErrorResponse{Error: err.Error()}
	WriteJSONResponse(statusCode, errorJSON, w, r)
}

func (h *inventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Add a new inventory item.
	var item models.InventoryItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	fmt.Println(item)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item is successfully created."))
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
