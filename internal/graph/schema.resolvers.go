package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.NewLogin) (*model.Login, error) {
	access, err := r.Resolver.AuthUserService.UserLogin(ctx, input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	accessMap, ok := access.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("access is not a map[string]interface{}: %v", access)
	}

	token, ok := accessMap["token"].(string)
	if !ok {
		return nil, fmt.Errorf("token is not a string: %v", accessMap["token"])
	}

	response := &model.Login{
		Token: token,
	}

	return response, nil
}

// CreateContactInfo is the resolver for the createContactInfo field.
func (r *mutationResolver) CreateContactInfo(ctx context.Context, input model.NewContactInfo) (*model.ContactInfo, error) {
	contact_info := database.ContactInfo{
		Name:  input.Name,
		Email: input.Email,
		Phone: sql.NullString{String: input.Phone, Valid: true},
	}

	create, err := r.Resolver.ContactInfoService.CreateContactInfo(&contact_info)

	if err != nil {
		return nil, err
	}

	response := &model.ContactInfo{
		ID:    create.ID,
		Name:  create.Name,
		Email: create.Email,
		Phone: create.Phone.String,
	}

	return response, nil
}

// CreatePreUser is the resolver for the createPreUser field.
func (r *mutationResolver) CreatePreUser(ctx context.Context, input model.NewPreUser) (*model.PreUser, error) {
	tenant, err := r.Resolver.RBCAService.CreateTenant(input.Name)

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Erro em Criar o vinculo do Usario",
		}
	}

	user := database.User{
		Name:     input.Name,
		Email:    input.Email,
		Status:   input.Status,
		TenantID: sql.NullString{String: tenant.ID, Valid: true},
	}

	created, err := r.Resolver.UserService.CreatePreUser(&user)

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Erro em criar o usuario",
		}
	}

	response := &model.PreUser{
		ID:    created.ID,
		Email: created.Email,
	}

	return response, nil
}

// CreateCompleteUser is the resolver for the createCompleteUser field.
func (r *mutationResolver) CreateCompleteUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := database.User{
		Phone:          sql.NullString{String: input.Phone, Valid: true},
		Password:       sql.NullString{String: input.Password, Valid: true},
		DocumentType:   sql.NullString{String: input.DocumentType, Valid: true},
		DocumentNumber: sql.NullString{String: input.DocumentNumber, Valid: true},
	}

	create, err := r.Resolver.UserService.CompleteRegisterUser(input.RegisterToken, &user)

	if err != nil {
		return nil, err
	}

	response := &model.User{
		ID:    create.ID,
		Email: create.Email,
	}

	return response, err
}

// CreateAddress is the resolver for the createAddress field.
func (r *mutationResolver) CreateAddress(ctx context.Context, input model.NewAddress) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: CreateAddress - createAddress"))
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input model.NewRole) (*model.Roles, error) {
	role := database.Role{
		Name:        input.Name,
		Description: input.Description,
	}

	created, err := r.Resolver.RBCAService.CreateRoles(&role)

	if err != nil {
		return nil, err
	}

	response := &model.Roles{
		ID:          created.ID,
		Name:        created.Name,
		Description: created.Description,
	}

	return response, nil
}

// CreatePermission is the resolver for the createPermission field.
func (r *mutationResolver) CreatePermission(ctx context.Context, input model.NewPermission) (*model.Permissions, error) {
	permission := database.Permission{
		Name:        input.Name,
		Description: input.Description,
	}

	created, err := r.Resolver.RBCAService.CreatePermissions(&permission)

	if err != nil {
		return nil, err
	}

	return &model.Permissions{
		ID:          created.ID,
		Name:        created.Name,
		Description: created.Description,
	}, nil
}

// CreateRolePermission is the resolver for the createRolePermission field.
func (r *mutationResolver) CreateRolePermission(ctx context.Context, input model.NewRolePermission) (*model.RolePermissions, error) {
	assign := database.RolesPermission{
		RoleID:       input.RoleID,
		PermissionID: input.PermissionID,
	}

	created, err := r.Resolver.RBCAService.CreateRolesPermissions(&assign)

	if err != nil {
		return nil, err
	}

	return &model.RolePermissions{
		RoleID:       created.RoleID,
		PermissionID: created.PermissionID,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	getUser, err := r.Resolver.UserService.GetUsers()

	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, user := range getUser {
		listUser := &model.User{
			ID:             user.ID,
			Name:           user.Name,
			Email:          user.Email,
			Phone:          user.Phone.String,
			DocumentNumber: user.DocumentNumber.String,
		}

		users = append(users, listUser)
	}

	return users, nil
}

// Address is the resolver for the address field.
func (r *queryResolver) Address(ctx context.Context, id string) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: Address - address"))
}

// Addresses is the resolver for the addresses field.
func (r *queryResolver) Addresses(ctx context.Context) ([]*model.Address, error) {
	panic(fmt.Errorf("not implemented: Addresses - addresses"))
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
