package repository

import (
	"context"
	"log"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type IImage interface {
	GetImageS3(id pgtype.UUID) (*database.TemplateEmail, error)
	UpdateImageUser(image *database.UpdateImageUserParams) (*pgtype.UUID, error)
	CreateImageUser(image *database.CreateImageUserParams) (*pgtype.UUID, error)
}

type Image struct {
	DB *database.Queries
}

func NewImageRepository(db *database.Queries) *Image {
	return &Image{
		DB: db,
	}
}

func (e *Image) GetImageS3(id pgtype.UUID) (*pgtype.Text, error) {
	get, err := e.DB.GetImageS3(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (e *Image) CreateImageUser(image *database.CreateImageUserParams) (*pgtype.UUID, error) {

	create, err := e.DB.CreateImageUser(context.Background(), database.CreateImageUserParams{
		ID:          image.ID,
		Url:         image.Url,
		Description: image.Description,
		CreatedAt:   image.CreatedAt,
		ID_2:        image.ID_2,
	})

	log.Println("Creating", create)

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (e *Image) GetImageUser(id pgtype.UUID) (*database.GetImageUserRow, error) {
	get, err := e.DB.GetImageUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (e *Image) UpdateImageUser(image *database.UpdateImageUserParams) (*pgtype.UUID, error) {
	log.Printf("Update", image)
	update, err := e.DB.UpdateImageUser(context.Background(), database.UpdateImageUserParams{
		Url:       image.Url,
		UpdatedAt: image.UpdatedAt,
		ID:        image.ID,
	})

	if err != nil {
		return nil, err
	}

	return &update, nil
}
