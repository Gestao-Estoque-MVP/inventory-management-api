package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type SMTP struct {
	Username string
	Password string
	Host     string
	Port     string
}

var smtpConf *SMTP

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

	smtpConf = &SMTP{
		Username: os.Getenv("MAIL_TRAP_USERNAME"),
		Password: os.Getenv("MAIL_TRAP_PASSWORD"),
		Host:     os.Getenv("MAIL_TRAP_HOST"),
		Port:     "587",
	}
}

func SendEmail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", smtpConf.Username, smtpConf.Password, smtpConf.Host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("From: " + "SwiftStock <mailtrap@swiftstock.com.br>" + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" + body)

	err := smtp.SendMail(smtpConf.Host+":"+smtpConf.Port, auth, os.Getenv("MAIL_TRAP_FROM"), to, msg)
	if err != nil {
		log.Printf("Erro ao enviar o email:  %v", err)
		return err
	}
	log.Printf("Email enviado com sucesso.")
	return nil
}

func SendEmailAsync(to []string, subject string, body string) {
	go func() {
		err := SendEmail(to, subject, body)
		if err != nil {
			log.Printf("Erro ao enviar o email assíncrono: %v", err)
		}
	}()
}
