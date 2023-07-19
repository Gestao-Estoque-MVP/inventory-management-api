package service

import (
	"database/sql"
	"errors"
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

	tenant, _ := us.role.CreateTenant(user.Name)

	params := &database.User{
		ID:             uuid.NewGen().NewV4().String(),
		Name:           user.Name,
		Email:          user.Email,
		Status:         user.Status,
		RoleID:         sql.NullString{String: assign.ID, Valid: true},
		TenantID:       sql.NullInt32{Int32: tenant.ID, Valid: true},
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

	verifyUser, err := us.userRepo.GetUser(id)

	if err != nil {
		log.Printf("Erro ao buscar usuário: %v\n", err)
	}

	if verifyUser.TokenExpiresAt.Time != time.Now().Add(1*time.Hour) {
		log.Printf("Sem permissão de criar usuário")
		return nil, errors.New("Sem permissão de criar usuário")
	}

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
