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
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	prod, err := produkService.GetByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // 404 kalau tidak ketemu
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": err.Error(), // "produk tidak ditemukan"
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// PUT /api/produk/{id}
func updateProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updatedProduk models.Produk
	if err := json.NewDecoder(r.Body).Decode(&updatedProduk); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	prod, err := produkService.Update(id, updatedProduk)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // 404 kalau tidak ketemu
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": err.Error(), // "produk tidak ditemukan"
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// DELETE /api/produk/{id}
func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	err = produkService.Delete(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // 404 kalau tidak ketemu
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": err.Error(), // "produk tidak ditemukan"
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sukses delete",
		"status":  "OK",
	})
}

// GET /api/produk
func getAllProduk(w http.ResponseWriter, r *http.Request) {
	prod, err := produkService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// POST /api/produk
func createProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru models.Produk
	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	prod, err := produkService.Create(produkBaru)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": err.Error(), // ambil langsung dari return error
		})
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
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
