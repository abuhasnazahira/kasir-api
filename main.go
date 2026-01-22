package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/bootstrap"
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

	// Welcome endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Selamat datang di Go Kasir",
		})
	})

	// Initial App
	bootstrap.InitApp()

	//initial and running server
	fmt.Print("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("gagal running server")
	}
}
