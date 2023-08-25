package service

import (
	"context"
	"log"

	"github.com/diogoX451/inventory-management-api/internal/graph/model"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
)

type AuthUser struct {
	us   repository.UserRepository
	role repository.RBCARepository
}

func NewAuthUser(us repository.UserRepository, role repository.RBCARepository) *AuthUser {
	return &AuthUser{
		us:   us,
		role: role,
	}
}

func (a *AuthUser) UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := a.us.GetUserByEmail(email)

	if err != nil {
		log.Printf("error em trazer user login: %v", err)
		return nil, err
	}

	role, err := a.role.GetRoleByID(getUser.RoleID.String)

	if err != nil {
		log.Printf("error em trazer user login: %v", err)
		return nil, err
	}

	if err := token.ComparePassword(getUser.Password.String, password); err != nil {
		log.Printf("Erro de comparar")
		return nil, err
	}

	jwtToken, err := JwtGenerate(ctx, getUser.ID, model.Role(role.Name))

	if err != nil {
		log.Printf("Erro em gerar o JWT")
		return nil, err
	}

	return map[string]interface{}{
		"token": jwtToken,
	}, nil
}
