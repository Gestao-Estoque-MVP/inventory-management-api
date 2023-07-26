package repository

import (
	"context"
	"database/sql"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type IAddressRepository interface {
	CreateAddress(address *database.Address) (*database.Address, error)
	UpdateAddress(*database.Address) (*database.Address, error)
	DeleteAddress(userID string) (*sql.Result, error)
	GetAddressByID(userID string) (database.Queries, error)
	GetAddress() ([]*database.Address, error)
}

type IAddress struct {
	DB *database.Queries
}

func NewAddressRepository(db *database.Queries) *IAddress {
	return &IAddress{
		DB: db,
	}
}

func (a *IAddress) CreateAddress(address *database.Address) (*database.Address, error) {
	create, err := a.DB.CreateAddress(context.Background(), database.CreateAddressParams{
		UserID:     address.UserID,
		Address:    address.Address,
		Number:     address.Number,
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		Country:    address.Country,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (a *IAddress) UpdateAddress(address *database.Address) (*database.Address, error) {
	update, err := a.DB.UpdateAddress(context.Background(), database.UpdateAddressParams{
		UserID:     address.UserID,
		Address:    address.Address,
		Number:     address.Number,
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		Country:    address.Country,
	})

	if err != nil {
		return nil, err
	}

	return &update, nil
}

func (a *IAddress) DeleteAddress(userID string) (*sql.Result, error) {
	delete, err := a.DB.DeleteAddress(context.Background(), userID)

	if err != nil {
		return nil, err
	}

	return &delete, nil
}

func (a *IAddress) GetAddressByID(userID string) (*database.Address, error) {
	getID, err := a.DB.GetAddressByID(context.Background(), userID)

	if err != nil {
		return nil, err
	}

	return &getID, nil
}

func (a *IAddress) GetAddress() ([]*database.Address, error) {
	list, err := a.DB.ListAddresses(context.Background())

	if err != nil {
		return nil, err
	}

	pointers := make([]*database.Address, len(list))

	for i := range list {
		address := list[i]
		pointers[i] = &address
	}

	return pointers, nil
}
