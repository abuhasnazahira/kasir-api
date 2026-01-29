package controllers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"kasir-api/services"
	"log"
	"net/http"
	"time"
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
	start := time.Now()
	log.Println("📥 Get Product By ID")

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

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"data": prod,
		},
	}

	log.Printf("✅ Success get product id=%d in %v\n", id, time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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
	start := time.Now()

	log.Println("📥 Incoming request:", r.Method, r.URL.Path)
	// Query all products with filtering, sorting, pagination if needed
	var req models.Pagination

	// decode JSON body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	// set nilai default jika tidak digunakan
	// Default values
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	log.Printf("🔎 Params → limit=%d offset=%d search='%s'\n",
		req.Limit, req.Offset, req.Search,
	)

	prod, totalRecords, totalFiltered, err := productService.GetAll(req.Limit, req.Offset, req.Search)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// FIX: ubah nil slice jadi empty slice
	if prod == nil {
		prod = []models.Product{}
	}

	resp := map[string]interface{}{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"payload": map[string]interface{}{
			"totalRecords":        totalRecords,
			"totalRecordFiltered": totalFiltered,
			"data":                prod,
		},
	}

	log.Printf("✅ Success → returned %d products in %v\n", len(prod), time.Since(start))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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
