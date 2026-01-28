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

var kategoriService *services.KategoriService

func InitKategoriController(svc *services.KategoriService) {
	kategoriService = svc
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
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	kat, err := kategoriService.GetByID(id)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// PUT /api/categories/{id}
func updateKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.KategoriByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	var updatedKategori models.Kategori
	if err := json.NewDecoder(r.Body).Decode(&updatedKategori); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	kat, err := kategoriService.Update(id, updatedKategori)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// DELETE /api/categories/{id}
func deleteKategoriByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDFromURL(r, routes.KategoriByID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Kategori ID")
		return
	}

	err = kategoriService.Delete(id)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessMessage(w, "Sukses delete")
}

// GET /api/categories
func getAllKategori(w http.ResponseWriter, r *http.Request) {
	kat, err := kategoriService.GetAll()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kat)
}

// POST /api/categories
func createKategori(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru models.Kategori
	if err := json.NewDecoder(r.Body).Decode(&kategoriBaru); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	prod, err := kategoriService.Create(kategoriBaru)
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
		switch r.Method {
		case http.MethodGet:
			getAllKategori(w, r)
		case http.MethodPost:
			createKategori(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
