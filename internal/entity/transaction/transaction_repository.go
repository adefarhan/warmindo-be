package transaction

import "gorm.io/gorm"

type TransactionRepository interface {
	CreateTransaction(transaction Transaction) error
}

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *GormTransactionRepository {
	return &GormTransactionRepository{db: db}
}

func (t *GormTransactionRepository) CreateTransaction(transaction Transaction) error {
	result := t.db.Create(&transaction)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
