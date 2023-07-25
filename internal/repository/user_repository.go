package repository

import (
	"context"
	"database/sql"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreatePreUser(*database.User) (*database.User, error)
	CreateCompleteUser(token string, user *database.User) (*database.CompleteRegisterUserRow, error)
	UpdateUser(id string, user *database.UpdateUserParams) error
	DeleteUser(id string) (*sql.Result, error)
	GetUser(id string) (*database.User, error)
	GetUsers() ([]*database.User, error)
	GetUserByEmail(email string) (*database.User, error)
	GetUserRegisterToken(token string) (*database.User, error)
}

type UserRepository struct {
	DB *database.Queries
}

func NewRepositoryUser(db *database.Queries) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (i *UserRepository) CreatePreUser(user *database.User) (*database.User, error) {
	create, err := i.DB.CreatePreRegisterUser(context.Background(), database.CreatePreRegisterUserParams{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		TenantID:       user.TenantID,
		RegisterToken:  user.RegisterToken,
		TokenExpiresAt: user.TokenExpiresAt,
		Status:         user.Status,
		RoleID:         user.RoleID,
		CreatedAt:      user.CreatedAt,
	})

	if err != nil {
		return nil, err
	}

	user = &database.User{
		ID:    create.ID,
		Name:  create.Name,
		Email: create.Email,
	}

	return user, nil
}

func (i *UserRepository) CreateCompleteUser(token string, user *database.User) (*database.CompleteRegisterUserRow, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	updateUser, err := i.DB.CompleteRegisterUser(context.Background(), database.CompleteRegisterUserParams{
		Phone:          user.Phone,
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
		Password:       sql.NullString{String: string(bytes), Valid: true},
		RegisterToken:  sql.NullString{String: token, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &database.CompleteRegisterUserRow{
		Name: updateUser.Name,
	}, nil
}

func (i *UserRepository) UpdateUser(id string, user *database.UpdateUserParams) error {

	err := i.DB.UpdateUser(context.Background(), database.UpdateUserParams{
		ID:             id,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
	})

	if err != nil {
		return err
	}

	return err
}

func (i *UserRepository) DeleteUser(id string) (*sql.Result, error) {
	delete, err := i.DB.DeleteUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &delete, nil
}

func (i *UserRepository) GetUser(id string) (*database.User, error) {
	get, err := i.DB.GetUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) GetUsers() ([]*database.User, error) {
	list, err := i.DB.ListUsers(context.Background())

	if err != nil {
		return nil, err
	}

	ptrList := make([]*database.User, len(list))
	for i := range list {
		user := list[i]
		ptrList[i] = &user
	}

	return ptrList, nil
}

func (i *UserRepository) GetUserByEmail(email string) (*database.User, error) {
	get, err := i.DB.GetUserByEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) GetUserRegisterToken(token string) (*database.User, error) {
	get, err := i.DB.GetUserRegisterToken(context.Background(), sql.NullString{String: token, Valid: true})

	if err != nil {
		return nil, err
	}

	return &get, nil
}
