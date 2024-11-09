package handler

import (
	"net/http"

	"hot-coffee/internal/service"
)

type InventoryHandler interface {
	AddInventoryItem(w http.ResponseWriter, r *http.Request)
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
