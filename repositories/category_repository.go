package repositories

import "kasir-api/models"

// CategoryRepository defines contract for CRUD operations
type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(p models.Category) (*models.Category, error)
	Update(id int, p models.Category) (*models.Category, error)
	Delete(id int) error
}
