package repository

import (
	"context"
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreatePreUser(*database.User) (*database.User, error)
	CreateCompleteUser(token string, user *database.User) (*database.CompleteRegisterUserRow, error)
	UpdateUser(id string, user *database.UpdateUserParams) error
	DeleteUser(id string) (bool, error)
	GetUser(id string) (*database.User, error)
	GetUsers() ([]*database.User, error)
	GetUserByEmail(email string) (*database.GetEmailRow, error)
	GetUserRegisterToken(token string) (*database.User, error)
	VerifyToken(token string) (*database.GetTokenPreRegisterRow, error)
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
		Password:       pgtype.Text{String: string(bytes), Valid: true},
		RegisterToken:  pgtype.Text{String: token, Valid: true},
		UpdatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true},
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

func (i *UserRepository) DeleteUser(id string) (bool, error) {
	_, err := i.DB.DeleteUser(context.Background(), id)

	if err != nil {
		return false, err
	}

	return true, nil
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
		log.Printf("Error getting users %v", err)
		return nil, err
	}

	ptrList := make([]*database.User, len(list))
	for i := range list {
		user := list[i]
		ptrList[i] = &user
	}

	return ptrList, nil
}

func (i *UserRepository) GetUserByEmail(email string) (*database.GetEmailRow, error) {
	get, err := i.DB.GetEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) GetUserRegisterToken(token string) (*database.User, error) {
	get, err := i.DB.GetUserRegisterToken(context.Background(), pgtype.Text{String: token, Valid: true})

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) VerifyToken(token string) (*database.GetTokenPreRegisterRow, error) {
	find, err := i.DB.GetTokenPreRegister(context.Background(), pgtype.Text{
		String: token,
		Valid:  true,
	})

	if err != nil {
		return nil, err
	}

	return &find, nil
}
