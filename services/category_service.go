package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAll(limit, offset int, search string) ([]models.Category, int, int, error) {
	return s.repo.GetAll(limit, offset, search)
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(p models.Category) (*models.Category, error) {
	return s.repo.Create(p)
}

func (s *CategoryService) Update(id int, p models.Category) (*models.Category, error) {
	return s.repo.Update(id, p)
}

func (s *CategoryService) Delete(id int) error {
	// cek apakah ada produk terkait kategori ini
	product, _ := s.repo.GetByCategoryID(id)
	if len(product) > 0 {
		return errors.New("Kategori masih digunakan oleh produk, tidak bisa dihapus")
	}

	return s.repo.Delete(id)
}
