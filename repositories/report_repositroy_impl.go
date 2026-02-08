package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type reportRepo struct {
	db *sql.DB
}

func NewReportRepository(database *sql.DB) ReportRepository {
	return &reportRepo{
		db: database,
	}
}

func (r *reportRepo) GetReport(start, end time.Time) (*models.ReportResponse, error) {
	var report models.ReportResponse

	// 1️⃣ Total revenue & total transaksi
	err := r.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, start, end).Scan(
		&report.TotalRevenue,
		&report.TotalTransaksi,
	)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Produk terlaris
	err = r.db.QueryRow(`
		SELECT 
			p."name",
			COALESCE(SUM(td.quantity), 0) AS qty
		FROM transaction_details td
		JOIN product p ON p.product_id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p."name"
		ORDER BY qty DESC
		LIMIT 1
	`, start, end).Scan(
		&report.ProdukTerlaris.Nama,
		&report.ProdukTerlaris.QtyTerjual,
	)

	// Tidak ada transaksi → bukan error
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.BestSeller{}
		return &report, nil
	}

	if err != nil {
		return nil, err
	}

	return &report, nil
}
