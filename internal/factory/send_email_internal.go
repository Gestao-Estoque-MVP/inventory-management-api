package factory

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/diogoX451/inventory-management-api/internal/service"
)

type SendEmailInternal struct {
	name           string
	email          string
	title          string
	templateID     string
	templateStruct map[string]interface{}
	S3             service.S3Service
}

func (si *SendEmailInternal) MultiEmail() (string, error) {
	return "", fmt.Errorf("Error")

}

func (si *SendEmailInternal) SendOneEmail() (string, error) {

	err := si.S3.GetTemplateObject(si.templateID)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	file, err := template.ParseFiles("./s3-template.html")
	if err != nil {
		return "", fmt.Errorf("Error get file: %v", err)
	}

	data := struct {
		Name  string
		Email string
	}{
		Name:  si.name,
		Email: si.email,
	}

	buffer := new(bytes.Buffer)
	if err = file.Execute(buffer, data); err != nil {
		return "", fmt.Errorf("Error executing file: % %", err)
	}

	send(buffer.String(), si.title)

	return "Send", nil

}

func send(body string, subject string) {}

func (s *SendEmail) SendOneEmail(email []string, name string) error {

	tmp, err := template.ParseFiles("./internal/templates/index.html")

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	if err = tmp.Execute(buf, data); err != nil {
		return err
	}

	s.body = buf.String()

	auth := smtp.PlainAuth("", os.Getenv("MAIL_TRAP_USERNAME"), os.Getenv("MAIL_TRAP_PASSWORD"), os.Getenv("MAIL_TRAP_HOST"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	s.subject = "Se prepare para inovação..."
	s.from = "SwiftStock <mailtrap@swiftstock.com.br>"
	msg := []byte("From: " + s.from + "\r\n" +
		"Subject: " + s.subject + "\r\n" +
		mime + "\r\n" + s.body)

	err = smtp.SendMail(os.Getenv("MAIL_TRAP_HOST")+":587", auth, os.Getenv("MAIL_TRAP_FROM"), email, msg)

	if err != nil {
		log.Printf("Erro ao enviar o email: %v", err)
		return err
	}

	log.Printf("Email enviado com sucesso ")

	return nil
}

func (s *SendEmail) SendMultiEmail(email []string) error {
	return nil
}

func findTemplate(name string) (string, error) {
	return "", fmt.Errorf("Rodando")
}
