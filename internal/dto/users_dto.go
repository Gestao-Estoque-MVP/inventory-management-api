package dto

import "github.com/jackc/pgx/v5/pgtype"

type UserCreateDTO struct {
	Name          string        `json:"name" binding:"required"`
	Email         string        `json:"email" binding:"required"`
	Password      string        `json:"password" binding:"required"`
	Document      string        `json:"document" binding:"required"`
	MobilePhone   string        `json:"mobile_phone" binding:"required"`
	CompanyId     pgtype.UUID   `json:"company_id"`
	RoleId        pgtype.UUID   `json:"role_id"`
	PermissionsId []pgtype.UUID `json:"permissions_id"`
}
