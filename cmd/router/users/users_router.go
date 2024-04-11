package users_router

import (
	"github.com/diogoX451/inventory-management-api/internal/database"
	users_handler "github.com/diogoX451/inventory-management-api/internal/handler/users"
	middlewares "github.com/diogoX451/inventory-management-api/internal/middleware"
	companies_repository "github.com/diogoX451/inventory-management-api/internal/repository/companies"
	rcba_repository "github.com/diogoX451/inventory-management-api/internal/repository/rcba"
	users_repository "github.com/diogoX451/inventory-management-api/internal/repository/users"
	users_services "github.com/diogoX451/inventory-management-api/internal/service/users"
	"github.com/gin-gonic/gin"
)

func RouterUsers(db *database.Queries, authorize *gin.RouterGroup, route *gin.RouterGroup) {
	user := users_services.NewUserCreateService(users_repository.NewRepositoryUsers(db), companies_repository.NewRepositoryCompanies(db), rcba_repository.NewRBCARepository(db))
	userHandler := users_handler.NewUsersHandler(user)

	route.POST("/login", userHandler.LoginUser)

	authorize.POST("/create-user", middlewares.Auth(), userHandler.CreateUser)
}