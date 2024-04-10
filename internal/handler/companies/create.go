package companies_handler

import (
	"github.com/diogoX451/inventory-management-api/internal/dto"
	companies_services "github.com/diogoX451/inventory-management-api/internal/service/companies"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ICompaniesHandler interface {
	CreateCompany(company dto.CreateCompanyDTO) (pgtype.UUID, error)
}

type CompaniesHandler struct {
	service companies_services.ICompaniesService
}

func NewCompaniesHandler(service companies_services.ICompaniesService) *CompaniesHandler {
	return &CompaniesHandler{service: service}
}

func (h *CompaniesHandler) CreateCompany(c *gin.Context) {
	var company dto.CreateCompanyDTO
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	create, err := h.service.CreateCompany(company)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": create})
}
