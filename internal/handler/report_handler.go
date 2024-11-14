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
	WriteInfoResponse(statusCode int, message string, w http.ResponseWriter, r *http.Request)
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
	if statusCode/100 >= 5 {
		h.logger.PrintErrorMsg(err.Error())
	} else {
		h.logger.PrintDebugMsg(err.Error())
	}

	errorJSON := &models.ErrorResponse{Error: err.Error()}
	h.WriteJSONResponse(statusCode, errorJSON, w, r)
}

func (h *reportHandler) WriteInfoResponse(statusCode int, message string, w http.ResponseWriter, r *http.Request) {
	h.logger.PrintDebugMsg(message)

	infoJSON := &models.InfoResponse{Message: message}
	h.WriteJSONResponse(statusCode, infoJSON, w, r)
}

func (h *reportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve total sales.

	totalSales, err := h.ReportService.GetTotalSales()
	if err != nil {
		h.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}

	totalSalesJSON := &models.TotalSales{TotalSales: totalSales}

	h.WriteJSONResponse(http.StatusOK, totalSalesJSON, w, r)
}

func (h *reportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve popular items.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving popular items."))
}
