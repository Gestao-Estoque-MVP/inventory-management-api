package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type SendEmail struct {
	from    string
	to      []string
	subject string
	body    string
}

func (s *SendEmail) SendOneEmail(email []string, name string) error {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

	tmp, err := template.ParseFiles("./internal/templates/index.html")
	data := struct {
		Name string
	}{
		Name: name,
	}

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
