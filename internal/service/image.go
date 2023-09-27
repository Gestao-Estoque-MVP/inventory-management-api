package service

import (
	"log"

	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type ImageService struct {
	template repository.Image
	s3       *S3Service
}

func NewImageService(image repository.Image, s3 *S3Service) *ImageService {
	return &ImageService{
		template: image,
		s3:       s3,
	}
}

func (s *ImageService) GetTemplate(id pgtype.UUID) (*string, error) {
	consult, err := s.template.GetTemplate(id)
	log.Printf("[DEBUG] getting template", consult)
	if err != nil {
		return nil, err
	}

	url, err := s.s3.GetUrlS3(*consult)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
