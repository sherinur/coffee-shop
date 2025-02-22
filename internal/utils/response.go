package utils

import (
	"encoding/json"
	"net/http"

	"coffee-shop/models"
)

// WriteRawJSONResponse writes a raw JSON response to the HTTP response writer with a given status code.
// It sets the Content-Type header to "application/json" and encodes the provided jsonResponse into the response body.
// If there is an error during encoding, it writes an error response with the internal server error status code.
func WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, err, w, r)
	}
}

// WriteJSONResponse writes a formatted JSON response to the HTTP response writer with a given status code.
// It sets the Content-Type header to "application/json" and formats the provided jsonResponse with indentation.
// If there is an error during JSON formatting, it writes an error response with the internal server error status code.
func WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	formattedJSON, err := json.MarshalIndent(jsonResponse, "", " ")
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	w.Write(formattedJSON)
}

// WriteErrorResponse writes an error response in JSON format to the HTTP response writer.
// It logs the error message based on the provided status code and returns a JSON object
// with the error message in the response body.
func WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
	// if statusCode/100 >= 5 {
	// 	logger.LOGGER.PrintErrorMsg(err.Error())
	// } else {
	// 	logger.LOGGER.PrintDebugMsg(err.Error())
	// }

	errorJSON := &models.ErrorResponse{Error: err.Error()}

	WriteJSONResponse(statusCode, errorJSON, w, r)
}

func WriteInfoResponse(statusCode int, message string, w http.ResponseWriter, r *http.Request) {
	// logger.LOGGER.PrintDebugMsg(message)

	infoJSON := &models.InfoResponse{Message: message}
	WriteJSONResponse(statusCode, infoJSON, w, r)
}
