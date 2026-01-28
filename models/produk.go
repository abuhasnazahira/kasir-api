package models

//data modeling menggunakan struct
type Produk struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Stock    int      `json:"stock"`
	Category Kategori `json:"category"`
}
