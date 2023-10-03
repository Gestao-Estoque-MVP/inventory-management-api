package service

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	configs3 "github.com/diogoX451/inventory-management-api/pkg/configS3"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type S3Service struct {
	S3Repository repository.S3Repository
	Bucket       string
	Key          string
	Region       string
}

func NewServiceS3(S3Repository repository.S3Repository, bucket string, key string, region string) *S3Service {
	return &S3Service{
		S3Repository: S3Repository,
		Bucket:       bucket,
		Key:          key,
		Region:       region,
	}
}

func (s *S3Service) UploadTemplateS3(file io.Reader, template database.TemplateEmail) (*database.CreateTemplateRow, error) {
	keyPath := filepath.Join("template_email", template.Name)
	find := configs3.S3Config()
	upload := manager.NewUploader(find)
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
	id, _ := uuid.NewV4()

	create, err := s.S3Repository.UploadTemplateS3(database.TemplateEmail{

		ID:          pgtype.UUID{Bytes: id, Valid: true},
		Name:        template.Name,
		Url:         keyPath,
		Description: template.Description,
	})

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *S3Service) UploadImageS3(file io.Reader, name string) (string, error) {
	keyPath := filepath.Join("image", name)
	find := configs3.S3Config()
	upload := manager.NewUploader(find)
	_, err := upload.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      &s.Bucket,
		Key:         aws.String(keyPath),
		Body:        file,
		ContentType: aws.String("image/jpeg, image/png"),
	})

	if err != nil {
		log.Println("Erro ao upload", err)
		return "", &gqlerror.Error{
			Message: "Error uploading template",
		}
	}

	return keyPath, nil
}

func (s *S3Service) GetUrlS3(path string) (string, error) {

	presignClient := s3.NewPresignClient(configs3.S3Config())

	consult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}, func(po *s3.PresignOptions) {
		po.Expires = time.Duration(15 * time.Minute)
	})
	if err != nil {
		return "", err
	}

	return consult.URL, nil

}

func (s *S3Service) GetTemplateObject(id [16]byte) error {
	find, err := s.S3Repository.GetTemplateUrlS3(id)

	if err != nil {
		return err
	}

	obj, err := configs3.S3Config().GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(find),
	})

	defer obj.Body.Close()

	tmp, err := os.CreateTemp("", "s3-template.html")

	defer tmp.Close()

	_, err = io.Copy(tmp, obj.Body)

	if err != nil {
		log.Printf("Error creating temporary")
		return err
	}

	return nil
}
