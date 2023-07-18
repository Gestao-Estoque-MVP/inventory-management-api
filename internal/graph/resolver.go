package graph

import (
	"context"
	"database/sql"
	"time"

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
