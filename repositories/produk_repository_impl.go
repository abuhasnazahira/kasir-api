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

func (r *produkRepo) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query(`SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description 
		FROM product p 
		JOIN category c ON p.category_id = c.category_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produks []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category.ID, &p.Category.Name, &p.Category.Description)
		if err != nil {
			return nil, err
		}
		produks = append(produks, p)
	}
	return produks, nil
}

func (r *produkRepo) GetByID(id int) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(`SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description 
		FROM product p 
		JOIN category c ON p.category_id = c.category_id 
		WHERE p.product_id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category.ID, &p.Category.Name, &p.Category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Product tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Create(p models.Product) (*models.Product, error) {
	err := r.db.QueryRow("INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING product_id",
		p.Name, p.Price, p.Stock, p.Category.ID).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *produkRepo) Update(id int, p models.Product) (*models.Product, error) {
	_, err := r.db.Exec("UPDATE product SET name = $1, price = $2, stock = $3, category_id = $4 WHERE product_id = $5",
		p.Name, p.Price, p.Stock, p.Category.ID, id)
	if err != nil {
		return nil, err
	}

	// Select updated product by id
	var updatedProduct models.Product
	err = r.db.QueryRow(`SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description 
		FROM product p 
		JOIN category c ON p.category_id = c.category_id 
		WHERE p.product_id = $1`, id).
		Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price, &updatedProduct.Stock,
			&updatedProduct.Category.ID, &updatedProduct.Category.Name, &updatedProduct.Category.Description)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
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

func (r *produkRepo) GetByKategoriID(categoryID int) ([]models.Product, error) {
	rows, err := r.db.Query(`SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description 
		FROM product p 
		JOIN category c ON p.category_id = c.category_id 
		WHERE p.category_id = $1`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produks []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category.ID, &p.Category.Name, &p.Category.Description)
		if err != nil {
			return nil, err
		}
		produks = append(produks, p)
	}
	return produks, nil
}
