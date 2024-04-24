package orderdetail

import "gorm.io/gorm"

type OrderDetailRepository interface {
	CreateOrderDetail(orderDetail OrderDetail) error
}

type GormOrderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) *GormOrderDetailRepository {
	return &GormOrderDetailRepository{db: db}
}

func (g *GormOrderDetailRepository) CreateOrderDetail(orderDetail OrderDetail) error {
	result := g.db.Create(&orderDetail)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
