package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type S3Repository interface {
	UploadTemplateS3(template database.TemplateEmail) (*database.CreateTemplateRow, error)
	GetTemplateUrlS3(id string) (string, error)
}

type IS3 struct {
	DB *database.Queries
}

func NewS3Repository(db *database.Queries) *IS3 {
	return &IS3{
		DB: db,
	}
}

func (s *IS3) UploadTemplateS3(template database.TemplateEmail) (*database.CreateTemplateRow, error) {
	create, err := s.DB.CreateTemplate(context.Background(), database.CreateTemplateParams{
		ID:          template.ID,
		Name:        template.Name,
		Url:         template.Url,
		Description: template.Description,
		CreatedAt:   template.CreatedAt,
		UpdatedAt:   template.UpdatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &create, err
}

func (s *IS3) GetTemplateUrlS3(id string) (string, error) {
	get, err := s.DB.GetTemplateS3(context.Background(), id)

	if err != nil {
		return "", err
	}

	return get, err
}
