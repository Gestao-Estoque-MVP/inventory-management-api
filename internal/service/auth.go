package service

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/repository"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
)

type AuthUser struct {
	us repository.UserRepository
}

func NewAuthUser(us repository.UserRepository) *AuthUser {
	return &AuthUser{
		us: us,
	}
}

func (a *AuthUser) UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := a.us.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if err := token.ComparePassword(getUser.Password.String, password); err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, getUser.ID)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}
