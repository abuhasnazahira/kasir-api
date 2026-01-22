package repositories

import (
	"errors"
	"kasir-api/models"
)

// Implementasi in-memory repository, tetap return error
type produkRepo struct {
	data []models.Produk
}

// NewProdukRepositoryMemory buat instance in-memory repository
func NewProdukRepository() ProdukRepository {
	return &produkRepo{
		data: []models.Produk{
			{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10, KategoriID: 1},
			{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40, KategoriID: 2},
		},
	}
}

func (r *produkRepo) GetAll() ([]models.Produk, error) {
	return r.data, nil
}

func (r *produkRepo) GetByID(id int) (*models.Produk, error) {
	for _, p := range r.data {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("produk tidak ditemukan")
}

func (r *produkRepo) Create(p models.Produk) (*models.Produk, error) {
	p.ID = len(r.data) + 1
	r.data = append(r.data, p)
	return &p, nil
}

func (r *produkRepo) Update(id int, p models.Produk) (*models.Produk, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			p.ID = id
			r.data[i] = p
			return &r.data[i], nil
		}
	}
	// return nil, service nanti bisa bikin "produk tidak ditemukan"
	return nil, errors.New("produk tidak ditemukan")
}

func (r *produkRepo) Delete(id int) error {
	for i, p := range r.data {
		if p.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("produk tidak ditemukan")
}

func (r *produkRepo) GetByKategoriID(kategoriID int) ([]models.Produk, error) {
	var hasil []models.Produk
	for i := range r.data {
		if r.data[i].KategoriID == kategoriID {
			hasil = append(hasil, r.data[i])
		}
	}
	return hasil, nil
}
