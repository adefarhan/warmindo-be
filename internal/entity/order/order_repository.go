package order

import "gorm.io/gorm"

type OrderRepository interface {
	GetOrders() ([]Order, error)
	GetOrder(orderId string) (Order, error)
	CreateOrder(order Order) error
	SaveOrder(order Order) error
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{db: db}
}

func (g *GormOrderRepository) GetOrders() ([]Order, error) {
	var orders []Order

	result := g.db.Where("deleted_at IS NULL").Find(&orders).Order("created_at desc")

	if result.Error != nil {
		return orders, result.Error
	}

	return orders, nil
}

func (g *GormOrderRepository) GetOrder(orderId string) (Order, error) {
	var order Order

	result := g.db.Where("id = ? AND deleted_at IS NULL", orderId).Find(&order)

	if result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

func (g *GormOrderRepository) CreateOrder(order Order) error {
	result := g.db.Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GormOrderRepository) SaveOrder(order Order) error {
	result := g.db.Save(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
