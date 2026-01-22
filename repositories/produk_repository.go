package repositories

import "kasir-api/models"

// ProdukRepository defines contract for CRUD operations
type ProdukRepository interface {
	GetAll() ([]models.Produk, error)
	GetByID(id int) (*models.Produk, error)
	Create(p models.Produk) (*models.Produk, error)
	Update(id int, p models.Produk) (*models.Produk, error)
	Delete(id int) error
	GetByKategoriID(kategoriID int) ([]models.Produk, error)
}
