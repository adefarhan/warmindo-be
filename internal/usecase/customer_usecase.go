package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/customer"
	"github.com/google/uuid"
)

type CustomerUseCase struct {
	repository customer.CustomerRepository
}

func NewCustomerUseCase(repository customer.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{repository: repository}
}

func (uc *CustomerUseCase) CreateCustomer(customer customer.Customer) (customer.Customer, error) {
	customer.ID = uuid.NewString()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = nil
	customer.DeletedAt = nil

	err := uc.repository.CreateCustomer(customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (uc *CustomerUseCase) GetCustomers() ([]customer.Customer, error) {
	customers, err := uc.repository.GetCustomers()
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (uc *CustomerUseCase) GetCustomer(customerId string) (customer.Customer, error) {
	customer, err := uc.repository.GetCustomer(customerId)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (uc *CustomerUseCase) UpdateCustomer(customerId string, request customer.Customer) (customer.Customer, error) {
	customer, err := uc.repository.GetCustomer(customerId)
	if err != nil {
		return customer, err
	}

	if customer.ID == "" {
		log.Printf("Customer with id %s not found", customerId)
		return customer, errors.New("customer not found")
	}

	customer.Name = request.Name
	customer.PhoneNumber = request.PhoneNumber
	customer.Address = request.Address
	timeNow := time.Now()
	customer.UpdatedAt = &timeNow

	err = uc.repository.SaveCustomer(customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (uc *CustomerUseCase) DeletCustomer(customerId string) (customer.Customer, error) {
	customer, err := uc.repository.GetCustomer(customerId)
	if err != nil {
		return customer, err
	}

	if customer.ID == "" {
		log.Printf("Customer with id %s not found", customerId)
		return customer, errors.New("customer not found")
	}

	timeNow := time.Now()
	customer.DeletedAt = &timeNow

	err = uc.repository.SaveCustomer(customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
