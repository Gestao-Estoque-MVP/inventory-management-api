package service

import "fmt"

type S3Service struct {
}

func NewServiceS3() *S3Service {
	return &S3Service{}
}

func (s *S3Service) UploadTemplate() (string, error) {
	return "", fmt.Errorf("Erro")
}


