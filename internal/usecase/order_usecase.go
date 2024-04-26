package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/adefarhan/warmindo-be/internal/entity/order"
	"github.com/google/uuid"
)

type OrderUseCase struct {
	repository         order.OrderRepository
	customerRepository customer.CustomerRepository
}

func NewOrderUseCase(repository order.OrderRepository, customerRepository customer.CustomerRepository) *OrderUseCase {
	return &OrderUseCase{repository: repository, customerRepository: customerRepository}
}

func (uc *OrderUseCase) CreateOrder(order order.Order) (order.Order, error) {
	customer, err := uc.customerRepository.GetCustomer(order.CustomerID)
	if err != nil {
		return order, err
	}

	if customer.ID == "" {
		log.Printf("Customer with id %s not found", order.CustomerID)
		return order, errors.New("customer not found")
	}

	order.ID = uuid.NewString()
	order.Status = "Created"
	order.CreatedAt = time.Now()
	order.UpdatedAt = nil
	order.DeletedAt = nil

	err = uc.repository.CreateOrder(order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (uc *OrderUseCase) GetOrders() ([]order.Order, error) {
	orders, err := uc.repository.GetOrders()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (uc *OrderUseCase) GetOrder(orderId string) (order.Order, error) {
	order, err := uc.repository.GetOrder(orderId)

	if err != nil {
		log.Printf("Failed to get order: %s", err.Error())
		return order, err
	}

	return order, nil
}

func (uc *OrderUseCase) FinishOrder(orderId string) (order.Order, error) {
	order, err := uc.repository.GetOrder(orderId)
	if err != nil {
		log.Printf("Failed to get order: %s", err.Error())
		return order, err
	}

	if order.ID == "" {
		log.Printf("order not found: %s", orderId)
		return order, err
	}

	order.Status = "Completed"
	timeNow := time.Now()
	order.UpdatedAt = &timeNow

	err = uc.repository.SaveOrder(order)
	if err != nil {
		log.Printf("Failed to save order: %s", err.Error())
		return order, err
	}

	return order, nil
}
