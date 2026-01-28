package controllers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"kasir-api/services"
	"net/http"
)

var productService *services.ProductService

func InitProductHandler(svc *services.ProductService) {
	productService = svc
}

/*
====================
HANDLERS
====================
*/

// GET /api/produk/{id}
func getProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProductByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	prod, err := productService.GetByID(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// PUT /api/produk/{id}
func updateProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProductByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	var updatedProduk models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduk); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	prod, err := productService.Update(id, updatedProduk)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// DELETE /api/produk/{id}
func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProductByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Produk ID")
		return
	}

	err = productService.Delete(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.SuccessMessage(w, "Sukses delete")
}

// GET /api/produk
func getAllProduct(w http.ResponseWriter, r *http.Request) {
	prod, err := productService.GetAll()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// POST /api/produk
func createProduct(w http.ResponseWriter, r *http.Request) {
	var produkBaru models.Product
	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	prod, err := productService.Create(produkBaru)
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

func RegisterProductRoutes() {
	// GET / PUT / DELETE by ID
	http.HandleFunc(routes.ProductByID, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProductByID(w, r)
		case http.MethodPut:
			updateProductByID(w, r)
		case http.MethodDelete:
			deleteProductByID(w, r)
		default:
			helpers.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

	// GET / POST
	http.HandleFunc(routes.Product, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllProduct(w, r)
		case http.MethodPost:
			createProduct(w, r)
		default:
			helpers.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})
}
