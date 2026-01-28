package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

// Implementasi in-memory repository, tetap return error
type produkRepo struct {
	// data []models.Produk
	db *sql.DB
}

// NewProdukRepositoryMemory buat instance in-memory repository
func NewProdukRepository(database *sql.DB) ProdukRepository {
	return &produkRepo{
		db: database,
	}
}

func (r *produkRepo) GetAll() ([]models.Produk, error) {
	rows, err := r.db.Query("SELECT product_id, name, price, stock, category_id FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produks []models.Produk
	for rows.Next() {
		var p models.Produk
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		produks = append(produks, p)
	}
	return produks, nil
}

func (r *produkRepo) GetByID(id int) (*models.Produk, error) {
	var p models.Produk
	err := r.db.QueryRow("SELECT product_id, name, price, stock, category_id FROM product WHERE product_id = $1", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Product tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Create(p models.Produk) (*models.Produk, error) {
	err := r.db.QueryRow("INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Price, p.Stock, p.CategoryID).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Update(id int, p models.Produk) (*models.Produk, error) {
	p.ID = id
	_, err := r.db.Exec("UPDATE product SET name = $1, price = $2, stock = $3, category_id = $4 WHERE product_id = $5",
		p.Name, p.Price, p.Stock, p.CategoryID, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Product tidak ditemukan")
	}
	return nil
}

func (r *produkRepo) GetByKategoriID(categoryID int) ([]models.Produk, error) {
	rows, err := r.db.Query("SELECT id, nama, harga, stok, category_id FROM produk WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produks []models.Produk
	for rows.Next() {
		var p models.Produk
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		produks = append(produks, p)
	}
	return produks, nil
}
