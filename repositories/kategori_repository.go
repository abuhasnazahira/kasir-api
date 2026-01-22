package repositories

import "kasir-api/models"

// KategoriRepository defines contract for CRUD operations
type KategoriRepository interface {
	GetAll() ([]models.Kategori, error)
	GetByID(id int) (*models.Kategori, error)
	Create(p models.Kategori) (*models.Kategori, error)
	Update(id int, p models.Kategori) (*models.Kategori, error)
	Delete(id int) error
}
