package repositories

import "github.com/diogoX451/inventory-management-api/internal/database"

type IAddressRepository interface {
	createAddress(*database.Address) (*database.Address, error)
	updateAddress(*database.Address) (*database.Address, error)
	deleteAddress(id int) (*database.Address, error)
	getAddressByID(id int) (*database.Address, error)
	getAddresses() ([]*database.Address, error)
}
