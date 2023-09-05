package service

import (
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"nullprogram.com/x/uuid"
)

type ContactInfoService struct {
	contactInfoRepository *repository.ContactInfoRepository
	email                 *EmailService
}

func NewContactInfoService(contactInfoRepository *repository.ContactInfoRepository, email *EmailService) *ContactInfoService {
	return &ContactInfoService{contactInfoRepository: contactInfoRepository, email: email}
}

func (s *ContactInfoService) CreateContactInfo(info *database.ContactInfo) (*database.ContactInfo, error) {

	params := &database.ContactInfo{
		ID:        uuid.NewGen().NewV4().String(),
		Name:      info.Name,
		Email:     info.Email,
		Phone:     info.Phone,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	create, err := s.contactInfoRepository.CreateContactInfo(params)

	if err != nil {
		log.Printf("Error creating contact info %v:", err)
		return nil, err
	}

	go func(email string) {
		detail := &EmailDetails{
			To:         []string{email},
			Subject:    "Pré-Cadastro no SwiftStock",
			TemplateID: "9de3e4bb-2c5d-4cfe-9b74-801e42d18769",
		}

		err = s.email.SendEmail(detail, "contact")
	}(create.Email)

	return create, nil
}
