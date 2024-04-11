package companies_repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type ICompaniesRepository interface {
	CreateCompany(company dto.CreateCompanyDTO) (*pgtype.UUID, error)
	AssociateUserToCompany(companyId pgtype.UUID, userId pgtype.UUID) error
}

type CompaniesRepository struct {
	db *database.Queries
}

func NewRepositoryCompanies(db *database.Queries) *CompaniesRepository {
	return &CompaniesRepository{
		db: db,
	}
}

func (r *CompaniesRepository) CreateCompany(company dto.CreateCompanyDTO) (*pgtype.UUID, error) {
	id, err := r.db.CreateCompanies(context.Background(), database.CreateCompaniesParams{
		Name:      pgtype.Text{String: company.Name, Valid: true},
		Document:  pgtype.Text{String: company.Document, Valid: true},
		AddressID: company.AddressId,
		IsAdmin:   pgtype.Bool{Bool: company.IsAdmin, Valid: true},
	})
	if err != nil {
		return &pgtype.UUID{}, err
	}

	return &id, nil
}

func (r *CompaniesRepository) AssociateUserToCompany(companyId pgtype.UUID, userId pgtype.UUID) error {
	_, err := r.db.AssociateUserCompany(context.Background(), database.AssociateUserCompanyParams{
		CompanyID: companyId,
		UserID:    userId,
	})

	if err != nil {
		return err
	}

	return nil
}
