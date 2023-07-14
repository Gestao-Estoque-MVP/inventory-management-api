package repositories

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreateUser(*database.User) (*database.User, error)
}

type UserRepository struct {
	DB *database.Queries
}

func NewRepositoryUser(db *database.Queries) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (i *UserRepository) CreateUser(user *database.User) (*database.User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	params := database.CreateUserParams{
		Name:           user.Name,
		Lastname:       user.Lastname,
		Email:          user.Email,
		Phone:          user.Phone,
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
		Password:       string(bytes),
	}

	createdUser, err := i.DB.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
