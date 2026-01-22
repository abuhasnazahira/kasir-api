package repositories

import (
	"errors"
	"kasir-api/models"
)

// Implementasi in-memory repository, tetap return error
type kategoriRepo struct {
	data []models.Kategori
}

// NewKategoriRepository buat instance in-memory repository
func NewKategoriRepository() KategoriRepository {
	return &kategoriRepo{
		data: []models.Kategori{
			{ID: 1, Nama: "Makanan"},
			{ID: 2, Nama: "Minuman"},
		},
	}
}

func (r *kategoriRepo) GetAll() ([]models.Kategori, error) {
	return r.data, nil
}

func (r *kategoriRepo) GetByID(id int) (*models.Kategori, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			return &r.data[i], nil // pointer ke slice asli
		}
	}
	return nil, errors.New("kategori tidak ditemukan")
}

func (r *kategoriRepo) Create(p models.Kategori) (*models.Kategori, error) {
	p.ID = len(r.data) + 1 // tetap seperti sebelumnya
	r.data = append(r.data, p)
	return &p, nil
}

func (r *kategoriRepo) Update(id int, p models.Kategori) (*models.Kategori, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			p.ID = id
			r.data[i] = p
			return &r.data[i], nil
		}
	}
	return nil, errors.New("kategori tidak ditemukan")
}

func (r *kategoriRepo) Delete(id int) error {
	for i, p := range r.data {
		if p.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("kategori tidak ditemukan")
}
