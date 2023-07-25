package service

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
	"nullprogram.com/x/uuid"
)

type UserService struct {
	userRepo repository.IUserRepository
	role     repository.IRBCA
}

func NewServiceUser(userRepo repository.IUserRepository, role repository.IRBCA) *UserService {
	return &UserService{userRepo: userRepo, role: role}
}

func (us *UserService) CreatePreUser(user *database.User) (*database.User, error) {
	token, _ := token.GeneratedToken()
	assign, _ := us.role.GetRole("users")

	params := &database.User{
		ID:             uuid.NewGen().NewV4().String(),
		Name:           user.Name,
		Email:          user.Email,
		Status:         user.Status,
		RoleID:         sql.NullString{String: assign.ID, Valid: true},
		TenantID:       user.TenantID,
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

func (us *UserService) CompleteRegisterUser(RegisterToken string, user *database.User) (*database.CompleteRegisterUserRow, error) {

	verifyUser, err := us.userRepo.GetUserRegisterToken(RegisterToken)

	if err != nil || verifyUser == nil {
		log.Printf("Erro ao buscar usuário: %v\n", err)
		return nil, fmt.Errorf("no user found with register token %s", RegisterToken)
	}

	updateUser, err := us.userRepo.CreateCompleteUser(verifyUser.RegisterToken.String, user)

	if err != nil {
		log.Printf("Erro ao criar usuário completo %v\n", err)
		return nil, err
	}

	return updateUser, nil
}

func (us *UserService) UpdateUser(id string, user *database.UpdateUserParams) (*database.UpdateUserParams, error) {
	update, err := us.userRepo.UpdateUser(id, user)

	if update != nil {
		log.Printf("Erro ao atualizar usuário completo %v\n", update)
		return nil, err
	}

	return update, nil
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

func (us *UserService) GetUserByEmail(email string) (*database.User, error) {
	get, err := us.userRepo.GetUserByEmail(email)

	if err != nil {
		log.Printf("Erro ao trazer usuário %v\n", err)
		return nil, err
	}

	return get, nil
}
