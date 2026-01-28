package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProdukService struct {
	repo         repositories.ProdukRepository
	kategoriRepo repositories.KategoriRepository // untuk validasi kategori
}

func NewProdukService(repo repositories.ProdukRepository, kategoriRepo repositories.KategoriRepository) *ProdukService {
	return &ProdukService{
		repo:         repo,
		kategoriRepo: kategoriRepo,
	}
}

// GetAll ambil semua produk
func (s *ProdukService) GetAll() ([]models.Produk, error) {
	return s.repo.GetAll()
}

// GetByID ambil produk berdasarkan ID
func (s *ProdukService) GetByID(id int) (*models.Produk, error) {
	return s.repo.GetByID(id)
}

// Create buat produk baru
func (s *ProdukService) Create(p models.Produk) (*models.Produk, error) {
	// validasi kategori
	if _, err := s.kategoriRepo.GetByID(p.Category.ID); err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	return s.repo.Create(p)
}

// Update produk berdasarkan ID
func (s *ProdukService) Update(id int, p models.Produk) (*models.Produk, error) {
	// validasi kategori
	if _, err := s.kategoriRepo.GetByID(p.Category.ID); err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	return s.repo.Update(id, p)
}

// Delete produk berdasarkan ID
func (s *ProdukService) Delete(id int) error {
	return s.repo.Delete(id)
}
