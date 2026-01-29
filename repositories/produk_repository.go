package repositories

import "kasir-api/models"

// ProdukRepository defines contract for CRUD operations
type ProdukRepository interface {
	GetAll(limit, offset int, search string) ([]models.Product, int, int, error)
	GetByID(id int) (*models.Product, error)
	Create(p models.Product) (*models.Product, error)
	Update(id int, p models.Product) (*models.Product, error)
	Delete(id int) error
	GetByKategoriID(kategoriID int) ([]models.Product, error)
}
