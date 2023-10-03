package repository

import (
	"context"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type S3Repository interface {
	UploadTemplateS3(template database.TemplateEmail) (*database.CreateTemplateRow, error)
	GetTemplateUrlS3(id [16]byte) (string, error)
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
		CreatedAt:   pgtype.Timestamp{Time: time.Now().UTC().Local(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC().Local(), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &create, err
}

func (s *IS3) GetTemplateUrlS3(id [16]byte) (string, error) {
	get, err := s.DB.GetTemplateS3(context.Background(), pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		return "", err
	}

	return get, err
}
