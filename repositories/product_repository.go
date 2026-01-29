package repositories

import "kasir-api/models"

// ProductRepository defines contract for CRUD operations
type ProductRepository interface {
	GetAll(limit, offset int, search string) ([]models.Product, int, int, error)
	GetByID(id int) (*models.Product, error)
	Create(p models.Product) (*models.Product, error)
	Update(id int, p models.Product) (*models.Product, error)
	Delete(id int) error
	GetByCategoryID(categoryID int) ([]models.Product, error)
}
