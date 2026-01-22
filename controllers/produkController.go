package controllers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/routes"
	"net/http"
)

/*
====================
DATA (GLOBAL)
====================
*/

var produk = []models.Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
}

/*
====================
HELPER
====================
*/

// func getIDFromURL(r *http.Request) (int, error) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
// 	return strconv.Atoi(idStr)
// }

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

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
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

	for i := range produk {
		if produk[i].ID == id {
			updatedProduk.ID = id
			produk[i] = updatedProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DELETE /api/produk/{id}
func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.ProdukByID)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses delete",
				"status":  "OK",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
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
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(produk)

		case http.MethodPost:
			var produkBaru models.Produk
			if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
				http.Error(w, "Invalid Request Body", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
