package graph

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/99designs/gqlgen/plugin"
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/diogoX451/inventory-management-api/internal/service"
)

type Resolver struct {
	UserService           *service.UserService
	UserRepository        repository.IUserRepository
	ContactInfoRepository repository.IContactInfoRepository
	ContactInfoService    *service.ContactInfoService
	RBCARepository        repository.IRBCA
	RBCAService           *service.RCBAService
	AuthUserService       *service.AuthUser
}

func (r *Resolver) Login(ctx context.Context, args struct {
	email    string
	password string
}) (interface{}, error) {
	login, err := r.AuthUserService.UserLogin(ctx, args.email, args.password)

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (r *Resolver) CreateRoles(ctx context.Context, args struct {
	ID          string
	Name        string
	Description string
}) (*database.Role, error) {
	role := &database.Role{
		ID:          args.ID,
		Name:        args.Name,
		Description: args.Name,
	}

	create, err := r.RBCAService.CreateRoles(role)

	if err != nil {
		return nil, err
	}

	return create, err
}

func (r *Resolver) CreatePermissions(ctx context.Context, args struct {
	ID          string
	Name        string
	Description string
}) (*database.Permission, error) {
	permission := &database.Permission{
		ID:          args.ID,
		Name:        args.Name,
		Description: args.Description,
	}

	create, err := r.RBCAService.CreatePermissions(permission)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (r *Resolver) CreateRolesPermissions(ctx context.Context, args struct {
	RoleID       string
	PermissionID string
}) (*database.RolesPermission, error) {
	assign := &database.RolesPermission{
		RoleID:       args.RoleID,
		PermissionID: args.PermissionID,
	}

	create, err := r.RBCAService.CreateRolesPermissions(assign)

	if err != nil {
		return nil, err
	}

	return create, nil

}

func (r *Resolver) CreateContactInfo(ctx context.Context, args struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}) (*database.ContactInfo, error) {
	contact_info := &database.ContactInfo{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		Phone:     sql.NullString{String: args.Phone, Valid: true},
		CreatedAt: args.CreatedAt,
	}

	create, err := r.ContactInfoService.CreateContactInfo(contact_info)

	if err != nil {
		return nil, err
	}

	return create, nil
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
		Status:         database.UserStatus(args.Status),
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

func (r *Resolver) CreateCompleteUser(ctx context.Context, id string, args struct {
	Phone          string
	Password       string
	DocumentType   string
	DocumentNumber string
}) (*database.CompleteRegisterUserRow, error) {
	user := &database.User{
		Phone:          sql.NullString{String: args.Phone},
		Password:       sql.NullString{String: args.Password},
		DocumentType:   sql.NullString{String: args.DocumentType},
		DocumentNumber: sql.NullString{String: args.DocumentNumber},
	}

	create, err := r.UserService.CompleteRegisterUser(id, user)

	if err != nil {
		return nil, err
	}

	return create, err
}

func (r *Resolver) UpdateUser(id string, args struct {
	Name           string
	Email          string
	DocumentNumber string
	DocumentType   string
	Phone          string
}) (*database.UpdateUserParams, error) {
	updates := &database.UpdateUserParams{
		Name:           args.Name,
		Email:          args.Email,
		DocumentType:   sql.NullString{String: args.DocumentType, Valid: true},
		DocumentNumber: sql.NullString{String: args.DocumentNumber, Valid: true},
		Phone:          sql.NullString{String: args.Phone, Valid: true},
	}

	update, err := r.UserService.UpdateUser(id, updates)

	if err != nil {
		return nil, err
	}

	return update, nil

}
