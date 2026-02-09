package handlers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"kasir-api/services"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
====================
Definition
====================
*/

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

/*
====================
HANDLERS
====================
*/
// GET /api/categories/{id}
func (h *CategoryHandler) getCategoryByID(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	log.Println("Incoming request:", r.Method, r.URL.Path)

	id, err := helpers.GetIDFromURL(r, routes.CategoryByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	cat, err := h.service.GetByID(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": cat,
		},
	}

	log.Printf("Success get category id=%d in %v\n", id, time.Since(start))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /api/categories/{id}
func (h *CategoryHandler) updateCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.CategoryByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	var updatedKategori models.Category
	if err := json.NewDecoder(r.Body).Decode(&updatedKategori); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	kat, err := h.service.Update(id, updatedKategori)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": kat,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /api/categories/{id}
func (h *CategoryHandler) deleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.CategoryByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessMessage(w, "Sukses delete")
}

// GET /api/categories
func (h *CategoryHandler) getAllCategory(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	log.Println("Incoming request:", r.Method, r.URL.Path)
	// Query all products with filtering, sorting, pagination if needed
	var req models.Pagination
	name := strings.TrimSpace(r.URL.Query().Get("name"))

	pageInt, err := helpers.GetQueryInt(r, "page", 10)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid page parameter")
		return
	}
	pageSizeInt, err := helpers.GetQueryInt(r, "pageSize", 0)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid pageSize parameter")
		return
	}
	req.Limit = pageInt
	req.Offset = (pageInt - 1) * pageSizeInt
	req.Search = name

	// set nilai default jika tidak digunakan
	// Default values
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	log.Printf("Params → limit=%d offset=%d search='%s'\n",
		req.Limit, req.Offset, req.Search,
	)

	cat, totalRecords, totalFiltered, err := h.service.GetAll(req.Limit, req.Offset, req.Search)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// FIX: ubah nil slice jadi empty slice
	if cat == nil {
		cat = []models.Category{}
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"totalRecords":        totalRecords,
			"totalRecordFiltered": totalFiltered,
			"data":                cat,
		},
	}

	log.Printf("Success → returned %d categories in %v\n", len(cat), time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /api/categories
func (h *CategoryHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	prod, err := h.service.Create(newCategory)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusCreated,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": prod,
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
func (h *CategoryHandler) HandleCategoryId(w http.ResponseWriter, r *http.Request) {
	// GET / PUT / DELETE by ID
	switch r.Method {
	case http.MethodGet:
		h.getCategoryByID(w, r)
	case http.MethodPut:
		h.updateCategoryByID(w, r)
	case http.MethodDelete:
		h.deleteCategoryByID(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	// GET / POST
	switch r.Method {
	case http.MethodGet:
		h.getAllCategory(w, r)
	case http.MethodPost:
		h.createCategory(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
