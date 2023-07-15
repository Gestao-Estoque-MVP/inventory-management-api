package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/graph/model"
)

// CreateContactInfo is the resolver for the createContactInfo field.
func (r *mutationResolver) CreateContactInfo(ctx context.Context, input model.NewContactInfo) (*model.ContactInfo, error) {
	panic(fmt.Errorf("not implemented: CreateContactInfo - createContactInfo"))
}

// CreatePreUser is the resolver for the createPreUser field.
func (r *mutationResolver) CreatePreUser(ctx context.Context, input model.NewPreUser) (*model.PreUser, error) {
	user := database.User{
		Name:  input.Name,
		Email: input.Email,
	}

	created, err := r.Resolver.UserService.CreatePreUser(&user)

	if err != nil {
		return nil, err
	}

	response := &model.PreUser{
		ID:    created.ID,
		Email: created.Email,
	}

	return response, nil
}

// CreateCompleteUser is the resolver for the createCompleteUser field.
func (r *mutationResolver) CreateCompleteUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateCompleteUser - createCompleteUser"))
}

// CreateAddress is the resolver for the createAddress field.
func (r *mutationResolver) CreateAddress(ctx context.Context, input model.NewAddress) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: CreateAddress - createAddress"))
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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
