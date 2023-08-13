package service

import (
	"context"
	"io"
	"log"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"nullprogram.com/x/uuid"
)

type S3Service struct {
	S3           *s3.Client
	S3Repository repository.S3Repository
	Bucket       string
	Key          string
	Region       string
}

func NewServiceS3(S3 *S3Service, S3Repository repository.S3Repository, bucket string, key string, region string) *S3Service {
	return &S3Service{
		S3:           S3.S3,
		S3Repository: S3Repository,
		Bucket:       bucket,
		Key:          key,
		Region:       region,
	}
}

func (s *S3Service) UploadTemplate(file io.Reader, template database.TemplateEmail) (*database.CreateTemplateRow, error) {
	keyPath := filepath.Join("template_email", template.Name)

	upload := manager.NewUploader(s.S3)
	_, err := upload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      &s.Bucket,
		Key:         aws.String(keyPath),
		Body:        file,
		ContentType: aws.String("text/html; charset=utf-8"),
	})

	if err != nil {
		log.Println("Erro ao upload", err)
		return nil, &gqlerror.Error{
			Message: "Error uploading template",
		}
	}

	create, err := s.S3Repository.UploadTemplate(database.TemplateEmail{
		ID:          uuid.NewGen().NewV4().String(),
		Name:        template.Name,
		Url:         keyPath,
		Description: template.Description,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	})

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *S3Service) GetTemplate(id string) (*database.TemplateEmail, error) {
	findTemplate, err := s.S3Repository.GetTemplate(id)

	if err != nil {
		log.Printf("Error getting template")
		return nil, err
	}

	presignClient := s3.NewPresignClient(s.S3)

	consult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(findTemplate),
	}, func(po *s3.PresignOptions) {
		po.Expires = time.Duration(15 * time.Minute)
	})
	if err != nil {
		return nil, err
	}

	return &database.TemplateEmail{
		Url: consult.URL,
	}, nil

}
