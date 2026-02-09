package handlers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
)

/*
====================
Definition
====================
*/
type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

/*
====================
Checkout Handler
====================
*/
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	transaction, err := h.service.Checkout(req.Items, true)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusCreated,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": transaction,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

/*
====================
Handle Functions
====================
*/
// multiple item apa aja, quantity nya
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		helpers.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
