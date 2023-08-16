package service

import (
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
)

type TemplateService struct {
	template repository.TemplateEmail
}

func NewTemplateService(template repository.TemplateEmail) *TemplateService {
	return &TemplateService{
		template: template,
	}
}

func (s *TemplateService) GetTemplate(id string) (*database.TemplateEmail, error) {
	consult, err := s.template.GetTemplate(id)

	if err != nil {
		return nil, err
	}

	return consult, nil
}
