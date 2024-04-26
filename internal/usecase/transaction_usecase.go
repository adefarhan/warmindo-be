package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/order"
	"github.com/adefarhan/warmindo-be/internal/entity/transaction"
	"github.com/google/uuid"
)

type TransactionUseCase struct {
	repository      transaction.TransactionRepository
	orderRepository order.OrderRepository
}

func NewTransactionUseCase(repository transaction.TransactionRepository, orderRepository order.OrderRepository) *TransactionUseCase {
	return &TransactionUseCase{repository: repository, orderRepository: orderRepository}
}

func (uc *TransactionUseCase) CreateTransaction(transaction transaction.Transaction) (transaction.Transaction, error) {
	order, err := uc.orderRepository.GetOrder(transaction.OrderID)
	if err != nil {
		log.Printf("Failed to get order: %s", err.Error())
		return transaction, err
	}

	if order.ID == "" {
		log.Printf("Order with id %s not found", transaction.OrderID)
		return transaction, errors.New("order not found")
	}

	transaction.ID = uuid.NewString()
	transaction.TotalPayment = order.TotalPrice
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = nil
	transaction.DeletedAt = nil

	err = uc.repository.CreateTransaction(transaction)
	if err != nil {
		log.Printf("Failed to create transaction: %s", err.Error())
		return transaction, err
	}

	order.Status = "Process"
	timeNow := time.Now()
	order.UpdatedAt = &timeNow

	err = uc.orderRepository.SaveOrder(order)
	if err != nil {
		log.Printf("Failed to save order: %s", err.Error())
		return transaction, err
	}

	return transaction, nil
}
