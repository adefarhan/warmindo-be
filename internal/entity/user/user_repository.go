package user

import "gorm.io/gorm"

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUser(userId string) (User, error)
	CreateUser(user User) error
	SaveUser(user User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

// GetUsers implements UserRepository.
func (g *GormUserRepository) GetUsers() ([]User, error) {
	var users []User

	result := g.db.Where("deleted_at IS NULL").Find(&users).Order("created_at desc")

	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

// GetUser implements UserRepository.
func (g *GormUserRepository) GetUser(userId string) (User, error) {
	var user User

	result := g.db.Where("id = ? AND deleted_at IS NULL", userId).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// CreateUser implements UserRepository.
func (g *GormUserRepository) CreateUser(user User) error {
	result := g.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// SaveUser implements UserRepository.
func (g *GormUserRepository) SaveUser(user User) error {
	result := g.db.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
