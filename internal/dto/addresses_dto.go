package dto

type AddressCreateDTO struct {
	Street     string `json:"street" binding:"required"`
	Number     string `json:"number" binding:"required"`
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
}

type AddressDTO struct {
	ID         uint   `json:"id"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
