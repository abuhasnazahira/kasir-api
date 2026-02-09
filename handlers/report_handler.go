package handlers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"time"
)

/*
====================
Definition
====================
*/

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetTodayReport()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": report,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	startDate := q.Get("start_date")
	endDate := q.Get("end_date")

	var (
		result *models.ReportResponse
		err    error
	)

	if startDate == "" || endDate == "" {
		result, err = h.service.GetTodayReport()
	} else {
		start, err1 := time.Parse("2006-01-02", startDate)
		end, err2 := time.Parse("2006-01-02", endDate)

		if err1 != nil || err2 != nil {
			helpers.Error(w, http.StatusBadRequest, "invalid date format (YYYY-MM-DD)")
			return
		}

		end = end.Add(24 * time.Hour)
		result, err = h.service.GetReportByRange(start, end)
	}

	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": result,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
