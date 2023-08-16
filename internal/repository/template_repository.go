package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type ITemplate interface {
	GetTemplate(id string) (*database.TemplateEmail, error)
}

type TemplateEmail struct {
	DB *database.Queries
}

func NewTemplateRepository(db *database.Queries) *TemplateEmail {
	return &TemplateEmail{
		DB: db,
	}
}

func (e *TemplateEmail) GetTemplate(id string) (*database.TemplateEmail, error) {
	get, err := e.DB.GetTemplate(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}
