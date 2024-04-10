package users_services

import (
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/dto"
	companies_repository "github.com/diogoX451/inventory-management-api/internal/repository/companies"
	rcba_repository "github.com/diogoX451/inventory-management-api/internal/repository/rcba"
	users_repository "github.com/diogoX451/inventory-management-api/internal/repository/users"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
	"github.com/diogoX451/inventory-management-api/pkg/helpers"
)

type IUserService interface {
	CreateUser(user dto.UserCreateDTO) (dto.Mesage, error)
}

type UserCreateService struct {
	repo    users_repository.IUsersRepository
	company companies_repository.ICompaniesRepository
	rcba    rcba_repository.IRBCA
}

func NewUserCreateService(repo users_repository.IUsersRepository, company companies_repository.ICompaniesRepository, rcba rcba_repository.IRBCA) *UserCreateService {
	return &UserCreateService{repo: repo, company: company, rcba: rcba}
}

func (s *UserCreateService) CreateUser(user dto.UserCreateDTO) (dto.Mesage, error) {
	user.Name = helpers.OnlyName(user.Name)
	user.Document = helpers.OnlyDocument(user.Document)
	user.MobilePhone = helpers.OnlyNumbers(user.MobilePhone)
	user.Password = token.HashPassword(user.Password)

	create, err := s.repo.CreateUser(user)
	if err != nil {
		return dto.Mesage{}, err
	}

	go func() {
		_ = s.company.AssociateUserToCompany(user.CompanyId, *create)
		_, _ = s.rcba.CreateUsersRoles(&database.UsersRole{
			UserID: *create,
			RoleID: user.RoleId,
		})

		for _, permission := range user.PermissionsId {
			_, _ = s.rcba.CreateUsersPermissions(&database.UsersPermission{
				UserID:       *create,
				PermissionID: permission,
			})
		}
	}()

	return dto.Mesage{
		Mesage: "Usuário criado com sucesso",
	}, nil
}
