package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"
	"log"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &database.User{
		Name:           input.Name,
		Lastname:       input.LastName,
		Email:          input.Email,
		Phone:          input.Phone,
		DocumentType:   input.DocumentType,
		DocumentNumber: input.DocumentNumber,
		Password:       input.Password,
	}

	created, err := r.Resolver.UserService.CreateUser(user)

	if err != nil {
		log.Printf("TESTE")
		return nil, err
	}

	response := &model.User{
		Name:           created.Name,
		LastName:       created.Lastname,
		Email:          created.Email,
		Phone:          created.Phone,
		DocumentType:   created.DocumentType,
		DocumentNumber: created.DocumentNumber,
	}

	return response, nil
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
	panic(fmt.Errorf("not implemented: Users - users"))
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
