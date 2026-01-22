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

var kategori = []models.Kategori{
	{ID: 1, Nama: "Indomie Godog"},
	{ID: 2, Nama: "Vit 1000ml"},
}

/*
====================
HANDLERS
====================
*/

// GET /api/categories/{id}
func getKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.KategoriByID)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

// PUT /api/categories/{id}
func updateKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.KategoriByID)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	var updatedKategori models.Kategori
	if err := json.NewDecoder(r.Body).Decode(&updatedKategori); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for i := range kategori {
		if kategori[i].ID == id {
			updatedKategori.ID = id
			kategori[i] = updatedKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedKategori)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DELETE /api/produk/{id}
func deleteKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.KategoriByID)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range kategori {
		if p.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)

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

func RegisterKategoriRoutes() {
	// GET / PUT / DELETE by ID
	http.HandleFunc(routes.KategoriByID, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getKategoriByID(w, r)
		case http.MethodPut:
			updateKategoriByID(w, r)
		case http.MethodDelete:
			deleteKategoriByID(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET / POST
	http.HandleFunc(routes.Kategori, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(kategori)

		case http.MethodPost:
			var kategoriBaru models.Kategori
			if err := json.NewDecoder(r.Body).Decode(&kategoriBaru); err != nil {
				http.Error(w, "Invalid Request Body", http.StatusBadRequest)
				return
			}

			kategoriBaru.ID = len(kategori) + 1
			kategori = append(kategori, kategoriBaru)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(kategoriBaru)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
