package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo         repositories.ProdukRepository
	kategoriRepo repositories.CategoryRepository // untuk validasi kategori
}

func NewProductService(repo repositories.ProdukRepository, kategoriRepo repositories.CategoryRepository) *ProductService {
	return &ProductService{
		repo:         repo,
		kategoriRepo: kategoriRepo,
	}
}

// GetAll ambil semua produk
func (s *ProductService) GetAll(limit, offset int, search string) ([]models.Product, int, int, error) {
	return s.repo.GetAll(limit, offset, search)
}

// GetByID ambil produk berdasarkan ID
func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// Create buat produk baru
func (s *ProductService) Create(p models.Product) (*models.Product, error) {
	// validasi kategori
	if _, err := s.kategoriRepo.GetByID(p.Category.ID); err != nil {
		return nil, errors.New("Kategori tidak ditemukan")
	}

	return s.repo.Create(p)
}

// Update produk berdasarkan ID
func (s *ProductService) Update(id int, p models.Product) (*models.Product, error) {
	// validasi kategori
	if _, err := s.kategoriRepo.GetByID(p.Category.ID); err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	return s.repo.Update(id, p)
}

// Delete produk berdasarkan ID
func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
