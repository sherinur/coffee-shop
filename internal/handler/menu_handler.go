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

type MenuHandler interface {
	AddMenuItem(w http.ResponseWriter, r *http.Request)
	GetMenuItems(w http.ResponseWriter, r *http.Request)
	GetMenuItem(w http.ResponseWriter, r *http.Request)
	UpdateMenuItem(w http.ResponseWriter, r *http.Request)
	DeleteMenuItem(w http.ResponseWriter, r *http.Request)
	WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request)
}

type menuHandler struct {
	MenuService service.MenuService
	logger      *logger.Logger
}

func NewMenuHandler(s service.MenuService, l *logger.Logger) *menuHandler {
	return &menuHandler{MenuService: s, logger: l}
}

func (h *menuHandler) WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
	}
}

func (h *menuHandler) WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	formattedJSON, err := json.MarshalIndent(jsonResponse, "", " ")
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	w.Write(formattedJSON)
}

func (h *menuHandler) WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
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

func (h *menuHandler) GetMenuItems(w http.ResponseWriter, r *http.Request) {
	data, err := h.MenuService.RetrieveMenuItems()
	if err != nil {
		switch err {
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	h.logger.PrintDebugMsg("Retrieved Menu items")

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *menuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	data, err := h.MenuService.RetrieveMenuItem(itemId)
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
