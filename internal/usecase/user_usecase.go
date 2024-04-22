package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/adefarhan/warmindo-be/internal/entity/user"
	"github.com/google/uuid"
)

type UserUseCase struct {
	repository user.UserRepository
}

func NewUserUseCase(repository user.UserRepository) *UserUseCase {
	return &UserUseCase{repository: repository}
}

func (uc *UserUseCase) CreateUser(user user.User) (user.User, error) {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = nil
	user.DeletedAt = nil

	err := uc.repository.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUsers() ([]user.User, error) {
	users, err := uc.repository.GetUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (uc *UserUseCase) GetUser(userId string) (user.User, error) {
	user, err := uc.repository.GetUser(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(userId string, request user.User) (user.User, error) {
	user, err := uc.repository.GetUser(userId)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		log.Printf("User with id %s not found", userId)
		return user, errors.New("user not found")
	}

	user.Name = request.Name
	user.PhoneNumber = request.PhoneNumber
	user.Address = request.Address
	timeNow := time.Now()
	user.UpdatedAt = &timeNow

	err = uc.repository.SaveUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) DeletUser(userId string) (user.User, error) {
	user, err := uc.repository.GetUser(userId)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		log.Printf("User with id %s not found", userId)
		return user, errors.New("user not found")
	}

	timeNow := time.Now()
	user.DeletedAt = &timeNow

	err = uc.repository.SaveUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
