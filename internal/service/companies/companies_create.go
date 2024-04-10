package companies_services

import (
	"github.com/diogoX451/inventory-management-api/internal/dto"
	companies_repository "github.com/diogoX451/inventory-management-api/internal/repository/companies"
	address_service "github.com/diogoX451/inventory-management-api/internal/service/address"
	"github.com/diogoX451/inventory-management-api/pkg/helpers"
	"github.com/jackc/pgx/v5/pgtype"
)

type ICompaniesService interface {
	CreateCompany(company dto.CreateCompanyDTO) (pgtype.UUID, error)
	AssociateCompanyToUser(companyId pgtype.UUID, userId pgtype.UUID) error
}

type CompaniesService struct {
	company companies_repository.ICompaniesRepository
	address address_service.IAddressService
}

func NewCompaniesCreateService(company companies_repository.ICompaniesRepository, address address_service.IAddressService) *CompaniesService {
	return &CompaniesService{company: company, address: address}
}

func (c *CompaniesService) CreateCompany(company dto.CreateCompanyDTO) (pgtype.UUID, error) {
	company.Name = helpers.OnlyName(company.Name)
	company.Document = helpers.OnlyDocument(company.Document)

	createAdress, err := c.address.CreateAddress(company.AddressCreateDTO)

	if err != nil {
		return pgtype.UUID{}, err
	}

	company.AddressId = *createAdress

	create, err := c.company.CreateCompany(company)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return *create, nil
}
