package models

//data modeling menggunakan struct
type Kategori struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
