package service

import (
	"bytes"
	"context"
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
	To         []string
	Subject    string
	TemplateID string
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
		return s.sendOneEmail(s.details.To, s.details.TemplateID)
	case "multi":
		return s.sendMultiEmail(s.details.To, s.details.TemplateID)
	case "contacts":
		return s.contacts(s.details.To, s.details.TemplateID)
	case "contact":
		return s.contact(s.details.To, s.details.TemplateID)
	default:
		return &gqlerror.Error{
			Message: "Don't know how to send %v " + typesSend,
		}
	}
}

func (e *EmailService) getTemplateObject(templateID string) (string, error) {
	dir := "./internal/templates"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Error creating directory: %v", err)
			return "", err
		}
	}

	find, err := e.template.GetTemplateUrlS3(templateID)
	if err != nil {
		log.Printf("Error getting template")
		return "", err
	}

	obj, err := configs3.S3Config().GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(find),
	})
	if err != nil {
		log.Printf("Error getting object from S3: %v", err)
		return "", err
	}
	defer obj.Body.Close()

	tmp, err := os.CreateTemp(dir, "s3-template-*.html")
	if err != nil {
		log.Printf("Error creating temporary file: %v", err)
		return "", err
	}
	defer tmp.Close()

	_, err = io.Copy(tmp, obj.Body)
	if err != nil {
		log.Printf("Error copying to temporary file: %v", err)
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

	if len(s.details.Subject) > 0 {
		subject = s.details.Subject
	}

	email.SendEmailAsync([]string{to[0]}, subject, buf.String())
	return nil
}
func (s *EmailService) sendMultiEmail(to []string, templateID string) error {
	path, err := s.getTemplateObject(templateID)
	if err != nil {
		return &gqlerror.Error{
			Message: "Error getting template object " + err.Error(),
		}
	}

	defer os.Remove(path)

	tmp, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	go func(t *template.Template) {
		find, _ := s.user.GetUsersByEmail()
		for _, e := range find {
			user, err := s.user.GetUserByEmail(*e)
			if err != nil {
				log.Printf("Error getting user by email %v", err)
				continue
			}

			data := struct {
				Name string
			}{
				Name: user.Name,
			}
			buf := new(bytes.Buffer)
			if err = t.Execute(buf, data); err != nil {
				log.Printf("Error executing template: %s", err)
				break
			}

			subject := ""
			if len(s.details.Subject) > 0 {
				subject = s.details.Subject
			}

			email.SendEmailAsync([]string{*e}, subject, buf.String())
		}
	}(tmp)

	return nil
}

func (con *EmailService) contact(to []string, templateID string) error {
	path, err := con.getTemplateObject(templateID)

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
	find, err := con.user.GetContact(to[0])
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

	if len(con.details.Subject) > 0 {
		subject = con.details.Subject
	}

	email.SendEmailAsync([]string{to[0]}, subject, buf.String())
	return nil
}

func (con *EmailService) contacts(to []string, templateID string) error {
	path, err := con.getTemplateObject(templateID)
	if err != nil {
		return &gqlerror.Error{
			Message: "Error getting template object " + err.Error(),
		}
	}

	defer os.Remove(path)

	tmp, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	go func(t *template.Template) {
		find, err := con.user.GetContacts()
		if err != nil {
			log.Printf("Error getting contacts %v", err)
			return
		}
		for _, e := range find {
			user, err := con.user.GetContact(*e)
			if err != nil {
				log.Printf("Error getting user by email %v", err)
				continue
			}

			data := struct {
				Name string
			}{
				Name: user.Name,
			}
			buf := new(bytes.Buffer)
			if err = t.Execute(buf, data); err != nil {
				log.Printf("Error executing template: %s", err)
				break
			}

			subject := ""
			if len(con.details.Subject) > 0 {
				subject = con.details.Subject
			}

			email.SendEmailAsync([]string{*e}, subject, buf.String())
		}
	}(tmp)

	return nil
}
