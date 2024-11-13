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

// AddMenuItem handles the HTTP request to add a new menu item.
// It processes the request body, validates the input, and calls the service layer to add the item.
func (h *menuHandler) AddMenuItem(w http.ResponseWriter, r *http.Request) {
	// Check if the request body is nil, which indicates an empty request body.
	// If it is nil, return a Bad Request (400) response with an appropriate error message.
	if r.Body == nil {
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	// Ensure that the request body is closed once the function exits.
	defer r.Body.Close()

	// Declare a variable to hold the decoded menu item.
	var item models.MenuItem
	// Decode the incoming JSON request body into the menu item struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		// If decoding fails, return a Bad Request (400) response with the error.
		h.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	// Log a debug message to indicate that a new menu item is being added.
	h.logger.PrintDebugMsg("Adding new menu item: %+v", item)

	// Call the MenuService to add the new menu item to the menu.
	err := h.MenuService.AddMenuItem(item)
	if err != nil {
		// Handle specific errors returned by the service layer.
		switch err {
		case service.ErrNotUniqueID:
			// If the menu item ID is not unique, return a Conflict (409) response.
			h.WriteErrorResponse(http.StatusConflict, err, w, r)
			return
		default:
			// For any other errors, return an Internal Server Error (500) response.
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// Log an info message to indicate that the menu item was successfully added.
	h.logger.PrintInfoMsg("Successfully added new menu item: %+v", item)

	// Return a Created (201) response along with the newly added menu item in the response body.
	h.WriteJSONResponse(http.StatusCreated, item, w, r)
}

// GetMenuItems handles the HTTP request to retrieve all menu items.
// It calls the service layer to fetch the data and returns it to the client.
func (h *menuHandler) GetMenuItems(w http.ResponseWriter, r *http.Request) {
	// Retrieve the menu items from the MenuService.
	data, err := h.MenuService.RetrieveMenuItems()
	if err != nil {
		// If an error occurs while fetching the menu items, return an Internal Server Error (500) response.
		switch err {
		default:
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// Log a debug message indicating that the menu items were retrieved successfully.
	h.logger.PrintDebugMsg("Retrieved Menu items")

	// Set the status code to OK (200) and write the retrieved menu items to the response.
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetMenuItem handles the HTTP request to retrieve a specific menu item by its ID.
// It checks if the item ID is valid, calls the service layer to fetch the menu item,
// and returns the result to the client. In case of errors, it responds with the appropriate error message.
func (h *menuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	// Retrieve the item ID from the request URL path.
	itemId := r.PathValue("id")

	// Check if the item ID is valid.
	if len(itemId) == 0 {
		// If the item ID is invalid, return a Bad Request (400) response with an error message.
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	// Call the MenuService to retrieve the menu item by its ID.
	data, err := h.MenuService.RetrieveMenuItem(itemId)
	if err != nil {
		// If an error occurs during retrieval, handle it based on the type of error.
		switch err {
		case service.ErrNoItem:
			// If no item is found with the provided ID, return a Not Found (404) response.
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			// For any other errors, return an Internal Server Error (500) response.
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// Log a debug message indicating that the menu item with the specified ID was successfully retrieved.
	h.logger.PrintDebugMsg("Retrieved menu item with ID: %s", itemId)

	// Set the response status to OK (200) and write the retrieved data to the response.
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		// If there's an error writing the response, log the error.
		h.logger.PrintErrorMsg("Failed to write response: %v", err)
	}
}

// UpdateMenuItem handles the HTTP request to update an existing menu item.
// It checks if the request body is valid, decodes the new menu item data, and
// calls the service layer to update the menu item. In case of errors, it responds
// with the appropriate HTTP status and error message.
func (h *menuHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	// Check if the request body is empty.
	if r.Body == nil {
		// If the request body is empty, return a Bad Request (400) response with an error message.
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("request body can not be empty"), w, r)
		return
	}
	// Ensure the request body is closed after handling.
	defer r.Body.Close()

	// Retrieve the item ID from the request URL path.
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		// If the item ID is invalid, return a Bad Request (400) response with an error message.
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	// Decode the incoming request body into a MenuItem model.
	var item models.MenuItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		// If decoding fails, return a Bad Request (400) response with the decoding error.
		h.WriteErrorResponse(http.StatusBadRequest, err, w, r)
		return
	}

	// Call the MenuService to update the menu item with the provided ID and data.
	err := h.MenuService.UpdateMenuItem(itemId, item)
	if err != nil {
		// Handle errors based on the specific type of error returned from the service.
		switch err {
		case service.ErrNoItem:
			// If the item is not found, return a Not Found (404) response.
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		case service.ErrNotUniqueID:
			// If the new item ID is not unique, return a Bad Request (400) response with the error.
			h.WriteErrorResponse(http.StatusBadRequest, fmt.Errorf("item with id '%s' not unique", itemId), w, r)
			return
		default:
			// For any other errors, return an Internal Server Error (500) response.
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// If the update is successful, return an OK (200) status response.
	w.WriteHeader(http.StatusOK)
}

// DeleteMenuItem handles the HTTP request to delete a menu item by its ID.
// It validates the item ID, calls the service layer to delete the item, and
// responds with the appropriate HTTP status and message.
func (h *menuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	// Retrieve the item ID from the request URL path.
	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		// If the item ID is invalid, return a Bad Request (400) response with an error message.
		h.WriteErrorResponse(http.StatusBadRequest, errors.New("identificator is not valid"), w, r)
		return
	}

	// Call the MenuService to delete the menu item with the provided ID.
	err := h.MenuService.DeleteMenuItem(itemId)
	if err != nil {
		// Handle errors based on the specific type of error returned from the service.
		switch err {
		case service.ErrNoItem:
			// If the item is not found, return a Not Found (404) response with the error message.
			h.WriteErrorResponse(http.StatusNotFound, fmt.Errorf("item with id '%s' not found", itemId), w, r)
			return
		default:
			// For any other errors, return an Internal Server Error (500) response.
			h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
			return
		}
	}

	// Log the successful deletion of the menu item.
	h.logger.PrintDebugMsg("Menu item with ID: %s successfully deleted", itemId)

	// Return a No Content (204) status, indicating the item was successfully deleted with no response body.
	w.WriteHeader(http.StatusNoContent)
}
