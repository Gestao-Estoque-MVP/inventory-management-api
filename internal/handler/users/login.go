package users_handler

import (
	"github.com/diogoX451/inventory-management-api/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) LoginUser(c *gin.Context) {
	var user dto.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	login, err := h.service.LoginUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": login.Token})
}
