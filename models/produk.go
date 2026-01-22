package models

//data modeling menggunakan struct
type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	KategoriID int    `json:"kategoriId"`
}
