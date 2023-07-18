package service

import "github.com/diogoX451/inventory-management-api/internal/repositories"

type RCBAService struct {
	rcba *repositories.IRBCA
}

func NewRCBAService(rcba *repositories.IRBCA) *RCBAService {
	return &RCBAService{rcba: rcba}
}
