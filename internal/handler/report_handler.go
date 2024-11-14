package handler

import (
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
	"hot-coffee/pkg/logger"
)

type ReportHandler interface {
	GetTotalSales(w http.ResponseWriter, r *http.Request)
	GetPopularItems(w http.ResponseWriter, r *http.Request)
}

type reportHandler struct {
	ReportService service.ReportService
	logger        *logger.Logger
}

func NewReportHandler(rs service.ReportService, l *logger.Logger) *reportHandler {
	return &reportHandler{ReportService: rs, logger: l}
}

func (h *reportHandler) WriteInfoResponse(statusCode int, message string, w http.ResponseWriter, r *http.Request) {
	h.logger.PrintDebugMsg(message)

	infoJSON := &models.InfoResponse{Message: message}
	utils.WriteJSONResponse(statusCode, infoJSON, w, r)
}

func (h *reportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	totalSales, err := h.ReportService.GetTotalSales()
	if err != nil {
		h.logger.PrintErrorMsg("Failed to get total sales: " + err.Error())
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w, r)
		return
	}
	h.logger.PrintDebugMsg("Successfully retrieved the total sales: ", totalSales)
	utils.WriteJSONResponse(http.StatusOK, totalSales, w, r)
}

func (h *reportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logic to Retrieve popular items.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("There will be Retrieving popular items."))
}
