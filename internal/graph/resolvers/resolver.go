package resolvers

import (
	"github.com/diogoX451/inventory-management-api/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService        *service.UserService
	ContactInfoService *service.ContactInfoService
	RBCAService        *service.RCBAService
	AuthUserService    *service.AuthUser
	AddressService     *service.AddressService
	EmailService       *service.EmailService
	S3Service          *service.S3Service
}
