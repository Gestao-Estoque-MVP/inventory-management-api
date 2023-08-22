package factory

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/diogoX451/inventory-management-api/internal/repository"
)

type ISendEmail interface {
	MultiEmail() (string, error)
	SendOneEmail() (string, error)
}

type SendEmailFactory struct {
	templateRepo *repository.TemplateEmail
	userRepo     *repository.UserRepository
}

const (
	internal = "internal"
	external = "external"
)

func NewSendEmailFactory(template *repository.TemplateEmail, userRepo *repository.UserRepository) *SendEmailFactory {
	return &SendEmailFactory{
		templateRepo: template,
		userRepo:     userRepo,
	}
}

func Send(body string, subject string, email []string) error {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_TRAP_USERNAME"), os.Getenv("MAIL_TRAP_PASSWORD"), os.Getenv("MAIL_TRAP_HOST"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	if subject != "" {
		subject = "Informações para você"
	}
	from := os.Getenv("MAIL_TRAP_FROM_TITLE")
	msg := []byte(
		"From: " + from + "\r\n" +
			"Subject: " + subject + "\r\n" +
			mime + "\r\n" + body)

	send := smtp.SendMail(os.Getenv("MAIL_TRAP_HOST")+":587", auth, os.Getenv("MAIL_TRAP_FROM"), email, msg)
	if send != nil {
		log.Printf("Error sending mail: %v", send)
		return send
	}
	log.Printf("Sent successfully")
	return nil
}

func (send *SendEmailFactory) SendEmail(typeSend string, templateID string, userID string, title string) (ISendEmail, error) {
	switch typeSend {
	case internal:
		return &SendEmailInternal{
			userId:     userID,
			templateID: templateID,
			title:      title,
			template:   send.templateRepo,
			user:       send.userRepo,
		}, nil
	case external:
		return &SendEmailExternal{}, nil
	default:
		return nil, fmt.Errorf("Tipo de email incorreto passado")
	}
}
