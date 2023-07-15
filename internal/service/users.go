package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
	"nullprogram.com/x/uuid"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewServiceUser(userRepo repositories.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) CreatePreUser(user *database.User) (*database.User, error) {

	token, _ := token.GeneratedToken()
	user.Status = "pre-user"

	params := &database.User{
		ID:             uuid.NewGen().NewV4().String(),
		Name:           user.Name,
		Email:          user.Email,
		Status:         user.Status,
		RegisterToken:  sql.NullString{String: token, Valid: true},
		TokenExpiresAt: sql.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true},
		CreatedAt:      time.Now().Local(),
	}

	createUser, err := us.userRepo.CreatePreUser(params)

	if err != nil {
		log.Printf("Erro ao criar usuário: %v\n", err)
		return nil, err
	}

	return createUser, nil
}

func (us *UserService) CompleteRegisterUser(id string, user *database.User) (*database.User, error) {

	updateUser, err := us.userRepo.CreateCompleteUser(id, user)

	if err != nil {
		log.Printf("Erro ao criar usuário completo %v\n", err)
		return nil, err
	}

	return updateUser, nil
}

func (us *UserService) UpdateUser(id string, user *database.User) error {
	update := us.userRepo.UpdateUser(id, user)

	if update != nil {
		log.Printf("Erro ao atualizar usuário completo %v\n", update)
		return nil
	}

	return update
}

func (us *UserService) GetUsers() ([]*database.User, error) {
	list, err := us.userRepo.GetUsers()

	if err != nil {
		log.Printf("Erro ao trazer a listar usuário completo %v\n", err)
		return nil, err
	}

	return list, nil
}

func (us *UserService) GetUser(id string) (*database.User, error) {
	get, err := us.userRepo.GetUser(id)

	if err != nil {
		log.Printf("Erro ao trazer usuário  %v\n", err)
		return nil, err
	}

	return get, nil
}
