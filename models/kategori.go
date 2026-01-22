package models

//data modeling menggunakan struct
type Kategori struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
