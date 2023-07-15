package graph

import (
	"context"
	"database/sql"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
	"github.com/diogoX451/inventory-management-api/internal/service"
)

type Resolver struct {
	UserService    *service.UserService
	UserRepository repositories.IUserRepository
}

func (r *Resolver) CreatePreUser(ctx context.Context, args struct {
	ID             string
	Name           string
	Email          string
	Status         string
	RegisterToken  string
	TokenExpiresAt time.Time
	CreatedAt      time.Time
}) (*database.User, error) {
	user := &database.User{
		ID:             args.ID,
		Name:           args.Name,
		Email:          args.Email,
		Status:         args.Status,
		RegisterToken:  sql.NullString{String: args.RegisterToken, Valid: true},
		TokenExpiresAt: sql.NullTime{Time: args.TokenExpiresAt, Valid: true},
		CreatedAt:      args.CreatedAt,
	}

	createUser, err := r.UserService.CreatePreUser(user)

	if err != nil {
		return nil, err
	}

	return createUser, nil
}
