package users_services

import (
	"github.com/diogoX451/inventory-management-api/internal/dto"
	"github.com/diogoX451/inventory-management-api/internal/service/auth"
	"github.com/diogoX451/inventory-management-api/pkg/helpers"
)

func (u *UserService) LoginUser(user dto.Login) (dto.LoginOutput, error) {
	user.Email = helpers.OnlyEmail(user.Email)

	data, err := u.repo.GetUserByEmail(user.Email)
	if err != nil {
		return dto.LoginOutput{}, err
	}

	verifyPassword := helpers.VerifyPassword(user.Password, data.Password)

	if !verifyPassword {
		return dto.LoginOutput{}, err
	}

	t := auth.NewJWT()
	token, err := t.GenerateToken(data.ID, data.Role)

	if err != nil {
		return dto.LoginOutput{}, err
	}

	return dto.LoginOutput{
		Token: token,
	}, nil
}
