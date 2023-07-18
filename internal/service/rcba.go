package service

import "github.com/diogoX451/inventory-management-api/internal/repository"

type RCBAService struct {
	rcba *repository.IRBCA
}

func NewRCBAService(rcba *repository.IRBCA) *RCBAService {
	return &RCBAService{rcba: rcba}
}
