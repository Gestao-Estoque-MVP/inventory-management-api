package service

import (
	"io"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ImageService struct {
	image repository.Image
	s3    *S3Service
}

func NewImageService(image repository.Image, s3 *S3Service) *ImageService {
	return &ImageService{
		image: image,
		s3:    s3,
	}
}

func (s *ImageService) GetImageS3(id pgtype.UUID) (*string, error) {
	consult, err := s.image.GetImageS3(id)
	if err != nil {
		return nil, err
	}

	url, err := s.s3.GetUrlS3(consult.String)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (s *ImageService) UpdateImageUser(id pgtype.UUID, image io.Reader) (*pgtype.UUID, error) {
	consult, _ := s.image.GetImageUser(id)

	params := database.UpdateImageUserParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).Local(), Valid: true},
	}

	if !consult.Url.Valid {
		url, err := s.s3.UploadImageS3(image, consult.Name.String)
		if err != nil {
			return nil, err
		}
		params.Url = pgtype.Text{String: url, Valid: true}
		generateID, _ := uuid.NewV4()
		return s.image.CreateImageUser(&database.CreateImageUserParams{
			ID:          pgtype.UUID{Bytes: generateID, Valid: true},
			Description: pgtype.Text{String: consult.Name.String, Valid: true},
			Url:         pgtype.Text{String: url, Valid: true},
			CreatedAt:   pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).Local(), Valid: true},
			ID_2:        id,
		})
	}

	url, err := s.s3.UploadImageS3(image, consult.Name.String)
	if err != nil {
		return nil, err
	}

	params.Url = pgtype.Text{String: url, Valid: true}

	return s.image.UpdateImageUser(&params)

}
