package repositories

import (
	"database/sql"
	"errors"
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

	// 1. Buat transaksi (total_amount sementara 0)
	var transactionID int
	err = tx.QueryRow(`
		INSERT INTO transactions (total_amount)
		VALUES (0)
		RETURNING id
	`).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	totalAmount := 0
	details := make([]models.TransactionDetail, 0, len(items))

	// 2. Loop item checkout
	for _, item := range items {
		var (
			name  string
			price int
			stock int
		)

		// Lock product (hindari race condition)
		err := tx.QueryRow(`
			SELECT name, price, stock
			FROM product
			WHERE product_id = $1
			FOR UPDATE
		`, item.ProductID).Scan(&name, &price, &stock)
		if err != nil {
			return nil, err
		}

		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than 0")
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("stock not enough for product %s", name)
		}

		subtotal := price * item.Quantity
		totalAmount += subtotal

		// Update stok
		_, err = tx.Exec(`
			UPDATE product
			SET stock = stock - $1
			WHERE product_id = $2
		`, item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		// Insert detail + ambil ID
		var detailID int
		err = tx.QueryRow(`
			INSERT INTO transaction_details
				(transaction_id, product_id, quantity, subtotal)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`,
			transactionID,
			item.ProductID,
			item.Quantity,
			subtotal,
		).Scan(&detailID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ID:            detailID,
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			ProductName:   name,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	// 3. Update total transaksi
	_, err = tx.Exec(`
		UPDATE transactions
		SET total_amount = $1
		WHERE id = $2
	`, totalAmount, transactionID)
	if err != nil {
		return nil, err
	}

	// 4. Commit
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	committed = true

	// 5. Return result
	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
