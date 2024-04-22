package customer

import "gorm.io/gorm"

type CustomerRepository interface {
	GetCustomers() ([]Customer, error)
	GetCustomer(customerId string) (Customer, error)
	CreateCustomer(customer Customer) error
	SaveCustomer(customer Customer) error
}

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *GormCustomerRepository {
	return &GormCustomerRepository{db: db}
}

// GetCustomer implements UserRepository.
func (g *GormCustomerRepository) GetCustomers() ([]Customer, error) {
	var customers []Customer

	result := g.db.Where("deleted_at IS NULL").Find(&customers).Order("created_at desc")

	if result.Error != nil {
		return customers, result.Error
	}

	return customers, nil
}

// GetCustomer implements UserRepository.
func (g *GormCustomerRepository) GetCustomer(customerId string) (Customer, error) {
	var customer Customer

	result := g.db.Where("id = ? AND deleted_at IS NULL", customerId).Find(&customer)

	if result.Error != nil {
		return customer, result.Error
	}

	return customer, nil
}

// CreateCustomer implements UserRepository.
func (g *GormCustomerRepository) CreateCustomer(customer Customer) error {
	result := g.db.Create(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// SaveCustomer implements UserRepository.
func (g *GormCustomerRepository) SaveCustomer(customer Customer) error {
	result := g.db.Save(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
