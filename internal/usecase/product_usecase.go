package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/product"
	"github.com/google/uuid"
)

type ProductUseCase struct {
	repository product.ProductRepository
}

func NewProductUseCase(repository product.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: repository}
}

func (uc *ProductUseCase) CreateProduct(product product.Product) (product.Product, error) {
	product.ID = uuid.NewString()
	product.CreatedAt = time.Now()
	product.UpdatedAt = nil
	product.DeletedAt = nil

	err := uc.repository.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) GetProducts() ([]product.Product, error) {
	products, err := uc.repository.GetProducts()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (uc *ProductUseCase) GetProduct(productId string) (product.Product, error) {
	product, err := uc.repository.GetProduct(productId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) UpdateProduct(productId string, request product.Product) (product.Product, error) {
	product, err := uc.repository.GetProduct(productId)
	if err != nil {
		return product, err
	}

	if product.ID == "" {
		log.Printf("Product with id %s not found", productId)
		return product, errors.New("product not found")
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock
	timeNow := time.Now()
	product.UpdatedAt = &timeNow

	err = uc.repository.SaveProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) DeleteProduct(productId string) (product.Product, error) {
	product, err := uc.repository.GetProduct(productId)
	if err != nil {
		return product, err
	}

	if product.ID == "" {
		log.Printf("Product with id %s not found", productId)
		return product, errors.New("product not found")
	}

	timeNow := time.Now()
	product.DeletedAt = &timeNow

	err = uc.repository.SaveProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}
