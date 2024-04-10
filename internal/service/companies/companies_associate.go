package companies_services

import "github.com/jackc/pgx/v5/pgtype"

func (c *CompaniesService) AssociateCompanyToUser(companyId pgtype.UUID, userId pgtype.UUID) error {
	err := c.company.AssociateUserToCompany(companyId, userId)
	if err != nil {
		return err
	}

	return nil
}
