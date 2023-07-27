package service

import (
	"database/sql"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
)

type AddressService struct {
	addressRepo *repository.IAddress
}

func NewAddressService(addressRepo *repository.IAddress) *AddressService {
	return &AddressService{addressRepo: addressRepo}
}

func (s *AddressService) CreateAddress(address *database.Address) (*database.Address, error) {
	create, err := s.addressRepo.CreateAddress(&database.Address{
		UserID:     address.UserID,
		Address:    address.Address,
		Street:     address.Street,
		City:       address.City,
		Country:    address.Country,
		PostalCode: address.PostalCode,
		State:      address.State,
		Number:     address.Number,
		CreatedAt:  time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *AddressService) UpdateAddress(address *database.Address) (*database.Address, error) {
	update, err := s.addressRepo.UpdateAddress(&database.Address{
		UserID:     address.UserID,
		Address:    address.Address,
		Street:     address.Street,
		City:       address.City,
		Country:    address.Country,
		PostalCode: address.PostalCode,
		State:      address.State,
		Number:     address.Number,
		UpdatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return update, nil
}

func (s *AddressService) DeleteAddress(userID string) (*sql.Result, error) {
	delete, err := s.addressRepo.DeleteAddress(userID)

	if err != nil {
		return nil, err
	}

	return delete, nil
}

func (s *AddressService) GetAddressByID(userID string) (*database.Address, error) {
	list, err := s.addressRepo.GetAddressByID(userID)

	if err != nil {
		return nil, err
	}

	return list, nil

}
