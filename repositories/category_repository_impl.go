package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

// Implementasi in-memory repository, tetap return error
type categoryRepo struct {
	// data []models.Category
	db *sql.DB
}

// NewCategoryRepository buat instance in-memory repository
func NewCategoryRepository(database *sql.DB) CategoryRepository {
	return &categoryRepo{
		// data: []models.Category{
		// 	{ID: 1, Name: "Makanan", Description: "Produk makanan"},
		// 	{ID: 2, Name: "Minuman", Description: "Produk minuman"},
		// },
		db: database,
	}
}

func (r *categoryRepo) GetAll(limit, offset int, search string) ([]models.Category, int, int, error) {
	// total semua data
	var totalRecords int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM category`).Scan(&totalRecords)
	if err != nil {
		return nil, 0, 0, err
	}

	// total setelah filter
	var totalFiltered int
	err = r.db.QueryRow(`
		SELECT COUNT(*) FROM category
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%')
	`, search).Scan(&totalFiltered)
	if err != nil {
		return nil, 0, 0, err
	}

	query := `SELECT category_id, name, description FROM category `
	query += `WHERE (name ILIKE '%' || $1 || '%') `
	query += `ORDER BY category_id LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, search, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var kategoris []models.Category
	for rows.Next() {
		var k models.Category
		if err := rows.Scan(&k.ID, &k.Name, &k.Description); err != nil {
			return nil, 0, 0, err
		}
		kategoris = append(kategoris, k)
	}
	return kategoris, totalRecords, totalFiltered, nil
}

func (r *categoryRepo) GetByID(id int) (*models.Category, error) {
	var k models.Category
	err := r.db.QueryRow("SELECT category_id, name, description FROM category WHERE category_id = $1", id).Scan(&k.ID, &k.Name, &k.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Kategori tidak ditemukan")
		}
		return nil, err
	}
	return &k, nil
}

func (r *categoryRepo) Create(p models.Category) (*models.Category, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO category (name, description) VALUES ($1, $2) RETURNING category_id", p.Name, p.Description).Scan(&id)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return &p, nil
}

func (r *categoryRepo) Update(id int, p models.Category) (*models.Category, error) {
	result, err := r.db.Exec("UPDATE category SET name = $1, description = $2 WHERE category_id = $3", p.Name, p.Description, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("Kategori tidak ditemukan")
	}
	p.ID = id
	return &p, nil
}

func (r *categoryRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM category WHERE category_id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("kategori tidak ditemukan")
	}
	return nil
}
