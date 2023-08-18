package factory

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/diogoX451/inventory-management-api/internal/service"
)

type SendEmailInternal struct {
	templateID string
	userId     string
	title      string
	template   *repository.TemplateEmail
	user       *repository.UserRepository
	S3         *service.S3Service
}

func (si *SendEmailInternal) MultiEmail() (string, error) {
	return "", fmt.Errorf("Error")

}

func (si *SendEmailInternal) SendOneEmail() (string, error) {

	tmp, err := si.template.DB.GetTemplate(context.Background(), si.templateID)
	if err != nil {
		return "", fmt.Errorf("Error getting template: %v", err)
	}

	err = si.S3.GetTemplateObject(tmp.Url)
	if err != nil {
		return "", fmt.Errorf("Error getting template: %v", err)
	}

	file, err := template.ParseFiles("./s3-template.html")
	if err != nil {
		return "", fmt.Errorf("Error get file: %v", err)
	}

	user, err := si.user.GetUser(si.userId)
	if err != nil {
		return "", fmt.Errorf("Error getting user: %v", err)
	}

	data := struct {
		Name  string
		Email string
	}{
		Name:  user.Name,
		Email: user.Email,
	}

	buffer := new(bytes.Buffer)
	if err = file.Execute(buffer, data); err != nil {
		return "", fmt.Errorf("Error executing file: % %", err)
	}

	if err := Send(buffer.String(), "", []string{si.title}); err != nil {
		return "", fmt.Errorf("Error sending: %v", err)
	}

	return "Send", nil

}
