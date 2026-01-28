package controllers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"kasir-api/services"
	"net/http"
)

/*
====================
DATA (GLOBAL)
====================
*/

var categoryService *services.CategoryService

func InitCategoryHandler(svc *services.CategoryService) {
	categoryService = svc
}

/*
====================
HANDLERS
====================
*/
// GET /api/categories/{id}
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.CategoryByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	kat, err := categoryService.GetByID(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// PUT /api/categories/{id}
func updateCategoryByID(w http.ResponseWriter, r *http.Request) {
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

	kat, err := categoryService.Update(id, updatedKategori)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// DELETE /api/categories/{id}
func deleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.CategoryByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	err = categoryService.Delete(id)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessMessage(w, "Sukses delete")
}

// GET /api/categories
func getAllCategory(w http.ResponseWriter, r *http.Request) {
	kat, err := categoryService.GetAll()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// POST /api/categories
func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	prod, err := categoryService.Create(newCategory)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prod)
}

/*
====================
ROUTING
====================
*/
func RegisterCategoryRoutes() {
	// GET / PUT / DELETE by ID
	http.HandleFunc(routes.CategoryByID, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCategoryByID(w, r)
		case http.MethodPut:
			updateCategoryByID(w, r)
		case http.MethodDelete:
			deleteCategoryByID(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET / POST
	http.HandleFunc(routes.Category, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllCategory(w, r)
		case http.MethodPost:
			createCategory(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
