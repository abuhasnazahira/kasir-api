package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type KategoriService struct {
	repo       repositories.KategoriRepository
	produkRepo repositories.ProdukRepository // untuk cek apakah ada produk terkait
}

func NewKategoriService(repo repositories.KategoriRepository, produkRepo repositories.ProdukRepository) *KategoriService {
	return &KategoriService{
		repo:       repo,
		produkRepo: produkRepo,
	}
}

func (s *KategoriService) GetAll() ([]models.Kategori, error) {
	return s.repo.GetAll()
}

func (s *KategoriService) GetByID(id int) (*models.Kategori, error) {
	return s.repo.GetByID(id)
}

func (s *KategoriService) Create(p models.Kategori) (*models.Kategori, error) {
	return s.repo.Create(p)
}

func (s *KategoriService) Update(id int, p models.Kategori) (*models.Kategori, error) {
	return s.repo.Update(id, p)
}

func (s *KategoriService) Delete(id int) error {
	// cek apakah ada produk terkait kategori ini
	produk, _ := s.produkRepo.GetByKategoriID(id)
	if len(produk) > 0 {
		return errors.New("kategori masih digunakan oleh produk, tidak bisa dihapus")
	}

	return s.repo.Delete(id)
}
