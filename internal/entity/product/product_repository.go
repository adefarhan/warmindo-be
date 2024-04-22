package product

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProduct(productId string) (Product, error)
	CreateProduct(product Product) error
	SaveProduct(product Product) error
}

type GormProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (g *GormProductRepository) GetProducts() ([]Product, error) {
	var products []Product

	result := g.db.Where("deleted_at IS NULL").Find(&products).Order("created_at desc")

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (g *GormProductRepository) GetProduct(productId string) (Product, error) {
	var product Product

	result := g.db.Where("id = ? AND deleted_at IS NULL", productId).Find(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

func (g *GormProductRepository) CreateProduct(product Product) error {
	result := g.db.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GormProductRepository) SaveProduct(product Product) error {
	result := g.db.Save(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
