package handler

import (
	"log/slog"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
)

type ReportHandler interface {
	GetTotalSales(w http.ResponseWriter, r *http.Request)
	GetPopularItems(w http.ResponseWriter, r *http.Request)
}

type reportHandler struct {
	ReportService service.ReportService
	log           *slog.Logger
}

func NewReportHandler(rs service.ReportService, l *slog.Logger) *reportHandler {
	return &reportHandler{ReportService: rs, log: l}
}

func (h *reportHandler) WriteInfoResponse(statusCode int, message string, w http.ResponseWriter, r *http.Request) {
	h.log.Debug(message)

	infoJSON := &models.InfoResponse{Message: message}
	utils.WriteJSONResponse(statusCode, infoJSON, w, r)
}

func (h *reportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	totalSales, err := h.ReportService.GetTotalSales()
	if err != nil {
		h.log.Error("Failed to get total sales: " + err.Error())
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}
	h.log.Debug("Successfully retrieved the total sales", slog.Any("TotalSales", totalSales))
	utils.WriteJSONResponse(http.StatusOK, totalSales, w, r)
}

func (h *reportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve popular items.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving popular items."))
}
