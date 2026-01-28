package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo       repositories.CategoryRepository
	produkRepo repositories.ProdukRepository // untuk cek apakah ada produk terkait
}

func NewCategoryService(repo repositories.CategoryRepository, produkRepo repositories.ProdukRepository) *CategoryService {
	return &CategoryService{
		repo:       repo,
		produkRepo: produkRepo,
	}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
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
	produk, _ := s.produkRepo.GetByKategoriID(id)
	if len(produk) > 0 {
		return errors.New("kategori masih digunakan oleh produk, tidak bisa dihapus")
	}

	return s.repo.Delete(id)
}
