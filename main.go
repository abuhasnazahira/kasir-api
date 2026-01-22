package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/controllers"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"
)

func main() {
	//health endpoint for checking
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //header
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		}) // response json
		// w.Write([]byte("Ok"))
	}) //localhost:8080/health

	// buat repository
	produkRepo := repositories.NewProdukRepository()
	kategoriRepo := repositories.NewKategoriRepository()

	// buat service
	produkService := services.NewProdukService(produkRepo, kategoriRepo)
	kategoriService := services.NewKategoriService(kategoriRepo, produkRepo)

	// inject service ke controller
	controllers.InitProdukController(produkService)
	controllers.InitKategoriController(kategoriService)

	//REGISTER ROUTES
	controllers.RegisterProdukRoutes()
	controllers.RegisterKategoriRoutes()

	//initial and running server
	fmt.Print("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("gagal running server")
	}
}
