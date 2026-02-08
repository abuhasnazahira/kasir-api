package repositories

import "kasir-api/models"

// TransactionRepository defines contract for transaction operations
type TransactionRepository interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
}
