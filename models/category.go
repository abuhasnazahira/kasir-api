package models

//data modeling menggunakan struct
type Category struct {
	ID          int    `json:"id" validate:"required,gt=0"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}
