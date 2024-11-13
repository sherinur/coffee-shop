package handler

import (
	"encoding/json"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/models"
	"hot-coffee/pkg/logger"
)

type ReportHandler interface {
	GetTotalSales(w http.ResponseWriter, r *http.Request)
	GetPopularItems(w http.ResponseWriter, r *http.Request)

	WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request)
	WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request)
}

type reportHandler struct {
	ReportService service.ReportService
	logger        *logger.Logger
}

func NewReportHandler(rs service.ReportService, l *logger.Logger) *reportHandler {
	return &reportHandler{ReportService: rs, logger: l}
}

func (h *reportHandler) WriteRawJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
	}
}

func (h *reportHandler) WriteJSONResponse(statusCode int, jsonResponse any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	formattedJSON, err := json.MarshalIndent(jsonResponse, "", " ")
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	w.Write(formattedJSON)
}

func (h *reportHandler) WriteErrorResponse(statusCode int, err error, w http.ResponseWriter, r *http.Request) {
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

func (h *reportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve total sales.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving total sales."))
}

func (h *reportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve popular items.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving popular items."))
}
