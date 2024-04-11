package dto

import "github.com/jackc/pgx/v5/pgtype"

type CreateCompanyDTO struct {
	Name             string           `json:"name" binding:"required"`
	AddressCreateDTO AddressCreateDTO `json:"address" binding:"required"`
	Document         string           `json:"document" binding:"required"`
	AddressId        pgtype.UUID      `json:"address_id"`
	IsAdmin          bool             `json:"is_admin"`
}

type CompanyDTO struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Address  AddressDTO `json:"address"`
	Document string     `json:"document"`
}
