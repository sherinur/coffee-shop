package handler

import (
	"net/http"

	"hot-coffee/internal/service"
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

func (h *inventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Add a new inventory item.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be INVENTORY ITEM CREATING"))
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
