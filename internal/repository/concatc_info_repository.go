package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type IContactInfoRepository interface {
	CreateContactInfo(*database.ContactInfo) (*database.ContactInfo, error)
}

type ContactInfoRepository struct {
	DB *database.Queries
}

func NewRepositoryContactInfo(db *database.Queries) *ContactInfoRepository {
	return &ContactInfoRepository{
		DB: db,
	}
}

func (repo *ContactInfoRepository) CreateContactInfo(info *database.ContactInfo) (*database.ContactInfo, error) {
	create, err := repo.DB.CreateContactInfo(context.Background(), database.CreateContactInfoParams{
		ID:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Phone:     info.Phone,
		CreatedAt: info.CreatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}
