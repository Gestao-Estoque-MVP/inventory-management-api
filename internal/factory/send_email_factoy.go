package factory

import (
	"fmt"

	"github.com/diogoX451/inventory-management-api/internal/repository"
)

type ISendEmail interface {
	MultiEmail() (string, error)
	SendOneEmail() (string, error)
}

type SendEmailFactory struct {
	templateRepo repository.TemplateEmail
	userRepo     repository.UserRepository
	s3           repository.S3Repository
}

const (
	internal = "internal"
	external = "external"
)

func NewSendEmailFactory(template repository.TemplateEmail, userRepo repository.UserRepository, s3 repository.S3Repository) *SendEmailFactory {
	return &SendEmailFactory{
		templateRepo: template,
		userRepo:     userRepo,
		s3:           s3,
	}
}

func SendEmail(typeSend string, email string, templateID string, filter interface{}) (ISendEmail, error) {
	switch typeSend {
	case internal:
		return &SendEmailInternal{
			typeTemplate:   templateID,
			templateStruct: filter,
			to:             userID,
		}, nil
	case external:
		return &SendEmailExternal{}, nil
	default:
		return nil, fmt.Errorf("Tipo de email incorreto passado")
	}
}
