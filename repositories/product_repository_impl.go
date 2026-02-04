package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

// Implementasi in-memory repository, tetap return error
type productRepo struct {
	// data []models.Product
	db *sql.DB
}

// NewProductRepository buat instance in-memory repository
func NewProductRepository(database *sql.DB) ProductRepository {
	return &productRepo{
		db: database,
	}
}

func (r *productRepo) GetAll(limit, offset int, search string) ([]models.Product, int, int, error) {
	// total semua data
	var totalRecords int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM product`).Scan(&totalRecords)
	if err != nil {
		return nil, 0, 0, err
	}

	// total setelah filter
	var totalFiltered int
	err = r.db.QueryRow(`
		SELECT COUNT(*) FROM product
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
	`, search).Scan(&totalFiltered)
	if err != nil {
		return nil, 0, 0, err
	}

	query := `SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description `
	query += `FROM product p JOIN category c ON p.category_id = c.category_id `
	query += `WHERE (p.name ILIKE '%' || $1 || '%') `
	query += `ORDER BY p.product_id LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, search, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var produks []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category.ID, &p.Category.Name, &p.Category.Description)
		if err != nil {
			return nil, 0, 0, err
		}
		produks = append(produks, p)
	}

	return produks, totalRecords, totalFiltered, nil
}

func (r *productRepo) GetByID(id int) (*models.Product, error) {
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

func (r *productRepo) Create(p models.Product) (*models.Product, error) {
	var productExist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM product WHERE name = $1)", p.Name).Scan(&productExist)
	if err != nil {
		return nil, err
	}
	if productExist {
		return nil, errors.New("Product already exists")
	}

	var categoryExists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM category WHERE category_id = $1)", p.Category.ID).Scan(&categoryExists)
	if err != nil {
		return nil, err
	}
	if !categoryExists {
		return nil, errors.New("Category not found")
	}

	var productID int
	err = r.db.QueryRow("INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING product_id",
		p.Name, p.Price, p.Stock, p.Category.ID).Scan(&productID)
	if err != nil {
		return nil, err
	}

	// Select created product by id with category data
	var createdProduct models.Product
	err = r.db.QueryRow(`SELECT p.product_id, p.name, p.price, p.stock, c.category_id, c.name, c.description 
		FROM product p 
		JOIN category c ON p.category_id = c.category_id 
		WHERE p.product_id = $1`, productID).
		Scan(&createdProduct.ID, &createdProduct.Name, &createdProduct.Price, &createdProduct.Stock,
			&createdProduct.Category.ID, &createdProduct.Category.Name, &createdProduct.Category.Description)
	if err != nil {
		return nil, err
	}

	return &createdProduct, nil
}

func (r *productRepo) Update(id int, p models.Product) (*models.Product, error) {
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

func (r *productRepo) Delete(id int) error {
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

func (r *productRepo) GetByCategoryID(categoryID int) ([]models.Product, error) {
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
