package controllers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"kasir-api/services"
	"net/http"
)

var produkService *services.ProdukService

func InitProdukController(svc *services.ProdukService) {
	produkService = svc
}

/*
====================
HANDLERS
====================
*/

// GET /api/produk/{id}
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	prod, err := produkService.GetByID(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// PUT /api/produk/{id}
func updateProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	var updatedProduk models.Produk
	if err := json.NewDecoder(r.Body).Decode(&updatedProduk); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	prod, err := produkService.Update(id, updatedProduk)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// DELETE /api/produk/{id}
func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	err = produkService.Delete(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.SuccessMessage(w, "Sukses delete")
}

// GET /api/produk
func getAllProduk(w http.ResponseWriter, r *http.Request) {
	prod, err := produkService.GetAll()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// POST /api/produk
func createProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru models.Produk
	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	prod, err := produkService.Create(produkBaru)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
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

func RegisterProdukRoutes() {
	// GET / PUT / DELETE by ID
	http.HandleFunc(routes.ProdukByID, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProdukByID(w, r)
		case http.MethodPut:
			updateProdukByID(w, r)
		case http.MethodDelete:
			deleteProdukByID(w, r)
		default:
			helpers.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

	// GET / POST
	http.HandleFunc(routes.Produk, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllProduk(w, r)
		case http.MethodPost:
			createProduk(w, r)
		default:
			helpers.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})
}
