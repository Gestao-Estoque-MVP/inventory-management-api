package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	configs3 "github.com/diogoX451/inventory-management-api/pkg/configS3"
	"github.com/diogoX451/inventory-management-api/pkg/email"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type EmailDetails struct {
	to         []string
	subject    string
	templateID string
}

type EmailService struct {
	template repository.S3Repository
	user     repository.UserRepository
	details  *EmailDetails
}

func NewServiceEmail(template repository.S3Repository, user repository.UserRepository) *EmailService {
	return &EmailService{
		template: template,
		user:     user,
	}
}

var subject string = "Notitificação para você"

func (s EmailService) SendEmail(details *EmailDetails, typesSend string) error {
	s.details = details
	switch typesSend {
	case "one":
		return s.sendOneEmail(s.details.to, s.details.templateID)
	case "multi":
		return s.sendMultiEmail(s.details.to, s.details.templateID)
	default:
		return fmt.Errorf("Don't know how to send")
	}
}

func (e *EmailService) getTemplateObject(templateID string) (string, error) {
	find, err := e.template.GetTemplateUrlS3(templateID)
	if err != nil {
		log.Printf("Error getting template")
		return "", err
	}

	obj, err := configs3.S3Config().GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(find),
	})

	defer obj.Body.Close()

	tmp, err := os.CreateTemp("./internal/templates", "s3-template-*.html")

	defer tmp.Close()

	_, err = io.Copy(tmp, obj.Body)

	if err != nil {
		log.Printf("Error creating temporary", err)
		return "", err
	}

	return tmp.Name(), nil

}

func (s *EmailService) sendOneEmail(to []string, templateID string) error {
	path, err := s.getTemplateObject(templateID)

	if err != nil {
		log.Printf("Error getting template object")
		return &gqlerror.Error{
			Message: "Error getting template object " + err.Error(),
		}
	}

	tmp, err := template.ParseFiles(path)
	defer os.Remove(path)

	if err != nil {
		log.Printf("Error parsing")
		return err
	}

	find, err := s.user.GetUserByEmail(to[0])
	if err != nil {
		return &gqlerror.Error{
			Message: "Error getting user " + err.Error(),
		}
	}
	data := struct {
		Name string
	}{
		Name: find.Name,
	}

	buf := new(bytes.Buffer)

	if err = tmp.Execute(buf, data); err != nil {
		return err
	}

	if len(s.details.subject) > 0 {
		subject = s.details.subject
	}

	email.SendEmailAsync([]string{to[0]}, subject, buf.String())
	return nil
}

func (s *EmailService) sendMultiEmail(to []string, templateID string) error {
	return nil
}
