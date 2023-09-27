package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type IAddressRepository interface {
	CreateAddress(address *database.Address) (*database.Address, error)
	UpdateAddress(*database.Address) (*database.Address, error)
	DeleteAddress(userID [16]byte) (*pgconn.CommandTag, error)
	GetAddressByID(userID [16]byte) (*database.Address, error)
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
		CreatedAt:  address.CreatedAt,
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

func (a *IAddress) DeleteAddress(userID [16]byte) (bool, error) {
	_, err := a.DB.DeleteAddress(context.Background(), pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *IAddress) GetAddressByID(userID [16]byte) (*database.Address, error) {
	getID, err := a.DB.GetAddressByID(context.Background(), pgtype.UUID{Bytes: userID, Valid: true})

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
