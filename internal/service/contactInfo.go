package service

import (
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
	"nullprogram.com/x/uuid"
)

type ContactInfoService struct {
	contactInfoRepository *repositories.ContactInfoRepository
}

func NewContactInfoService(contactInfoRepository *repositories.ContactInfoRepository) *ContactInfoService {
	return &ContactInfoService{contactInfoRepository: contactInfoRepository}
}

func (s *ContactInfoService) CreateContactInfo(info *database.ContactInfo) (*database.ContactInfo, error) {

	params := &database.ContactInfo{
		ID:        uuid.NewGen().NewV4().String(),
		Name:      info.Name,
		Email:     info.Email,
		Phone:     info.Phone,
		CreatedAt: time.Now(),
	}

	create, err := s.contactInfoRepository.CreateContactInfo(params)

	if err != nil {
		log.Printf("Error creating contact info %v:", err)
		return nil, err
	}

	return create, nil
}
