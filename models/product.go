package models

//data modeling menggunakan struct
type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name" validate:"required,min=5"`
	Price    float64  `json:"price" validate:"required,gt=0"`
	Stock    int      `json:"stock" validate:"required,gt=0"`
	Category Category `json:"category" validate:"required"`
}
