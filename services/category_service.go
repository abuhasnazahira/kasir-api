package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo       repositories.CategoryRepository
	produkRepo repositories.ProductRepository // untuk cek apakah ada produk terkait
}

func NewCategoryService(repo repositories.CategoryRepository, produkRepo repositories.ProductRepository) *CategoryService {
	return &CategoryService{
		repo:       repo,
		produkRepo: produkRepo,
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
	produk, _ := s.produkRepo.GetByCategoryID(id)
	if len(produk) > 0 {
		return errors.New("kategori masih digunakan oleh produk, tidak bisa dihapus")
	}

	return s.repo.Delete(id)
}
