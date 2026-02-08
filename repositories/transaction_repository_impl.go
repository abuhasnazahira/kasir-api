package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepository(database *sql.DB) TransactionRepository {
	return &transactionRepo{
		db: database,
	}
}

func (repo *transactionRepo) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	var transactionID int
	if err := tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES (0) RETURNING id",
	).Scan(&transactionID); err != nil {
		return nil, err
	}

	totalAmount := 0
	details := make([]models.TransactionDetail, 0, len(items))

	for _, item := range items {
		var name string
		var price, stock int

		if err := tx.QueryRow(
			"SELECT name, price, stock FROM product WHERE product_id = $1 FOR UPDATE",
			item.ProductID,
		).Scan(&name, &price, &stock); err != nil {
			return nil, err
		}

		if item.Quantity <= 0 || stock < item.Quantity {
			return nil, fmt.Errorf("invalid stock for product %s", name)
		}

		subtotal := price * item.Quantity
		totalAmount += subtotal

		if _, err := tx.Exec(
			"UPDATE product SET stock = stock - $1 WHERE product_id = $2",
			item.Quantity, item.ProductID,
		); err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			ProductName:   name,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	if _, err := tx.Exec(
		"UPDATE transactions SET total_amount = $1 WHERE id = $2",
		totalAmount, transactionID,
	); err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO transaction_details
		(transaction_id, product_id, quantity, subtotal)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, d := range details {
		if _, err := stmt.Exec(
			d.TransactionID,
			d.ProductID,
			d.Quantity,
			d.Subtotal,
		); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	committed = true

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
