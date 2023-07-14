package graph

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
	"github.com/diogoX451/inventory-management-api/internal/service"
)

type Resolver struct {
	UserService    *service.UserService
	UserRepository repositories.IUserRepository
}

func (r *Resolver) CreateUser(ctx context.Context, args struct {
	ID             string
	Name           string
	Lastname       string
	Email          string
	Phone          string
	DocumentType   string
	DocumentNumber string
	Password       string
}) (*database.User, error) {
	user := &database.User{
		Name:           args.Name,
		Lastname:       args.Lastname,
		Email:          args.Email,
		Phone:          args.Phone,
		DocumentType:   args.DocumentType,
		DocumentNumber: args.DocumentNumber,
		Password:       args.Password,
	}

	createUser, err := r.UserService.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return createUser, nil
}
