package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type IImage interface {
	GetTemplate(id pgtype.UUID) (*database.TemplateEmail, error)
}

type Image struct {
	DB *database.Queries
}

func NewImageRepository(db *database.Queries) *Image {
	return &Image{
		DB: db,
	}
}

func (e *Image) GetTemplate(id pgtype.UUID) (*string, error) {
	get, err := e.DB.GetImageS3(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}
