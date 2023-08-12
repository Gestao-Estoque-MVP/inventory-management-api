package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	S3 *s3.Client
}

func NewServiceS3(S3 *S3Service) *S3Service {
	return &S3Service{
		S3: S3.S3,
	}
}

func (s *S3Service) UploadTemplate() (string, error) {
	return "", fmt.Errorf("Erro")
}
