package users_handler

import (
	"github.com/diogoX451/inventory-management-api/internal/dto"
	users_services "github.com/diogoX451/inventory-management-api/internal/service/users"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type IUsersHandler interface {
	CreateUser(user dto.UserCreateDTO) (pgtype.UUID, error)
}

type UsersHandler struct {
	service users_services.IUserService
}

func NewUsersHandler(service users_services.IUserService) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) CreateUser(c *gin.Context) {
	var user dto.UserCreateDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	create, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": create})
}
