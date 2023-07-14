package service

import (
	"log"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewServiceUser(userRepo repositories.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) CreateUser(user *database.User) (*database.User, error) {
	createUser, err := us.userRepo.CreateUser(user)

	if err != nil {
		log.Printf("Erro ao criar usuário: %v\n", err)
		return nil, err
	}

	return createUser, nil
}
