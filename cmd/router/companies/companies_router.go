package companies_router

import (
	"github.com/diogoX451/inventory-management-api/internal/database"
	companies_handler "github.com/diogoX451/inventory-management-api/internal/handler/companies"
	middlewares "github.com/diogoX451/inventory-management-api/internal/middleware"
	address_repository "github.com/diogoX451/inventory-management-api/internal/repository/address"
	companies_repository "github.com/diogoX451/inventory-management-api/internal/repository/companies"
	address_service "github.com/diogoX451/inventory-management-api/internal/service/address"
	companies_services "github.com/diogoX451/inventory-management-api/internal/service/companies"
	"github.com/gin-gonic/gin"
)

func RouterCompanies(db *database.Queries, route *gin.RouterGroup) {
	address := address_service.NewAddressService(address_repository.NewAddressRepository(db))
	company := companies_services.NewCompaniesCreateService(companies_repository.NewRepositoryCompanies(db), address)
	companyHandler := companies_handler.NewCompaniesHandler(company)

	route.POST("/create-company",
		middlewares.Auth(),
		middlewares.AuthSuperAdmin(),
		companyHandler.CreateCompany,
	)
}
