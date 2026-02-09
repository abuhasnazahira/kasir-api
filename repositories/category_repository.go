package repositories

import "kasir-api/models"

// CategoryRepository defines contract for CRUD operations
type CategoryRepository interface {
	GetAll(limit, offset int, search string) ([]models.Category, int, int, error)
	GetByID(id int) (*models.Category, error)
	Create(p models.Category) (*models.Category, error)
	Update(id int, p models.Category) (*models.Category, error)
	Delete(id int) error
	GetByCategoryID(categoryID int) ([]models.Product, error)
}
