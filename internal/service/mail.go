package service

import (
	"log"
	"net/smtp"
)

type SendEmail struct {
	to   []string
	from string
	msg  []byte
}

func (s *SendEmail) Send(email []string) error {
	auth := smtp.PlainAuth(
		"",
		"d77214b7e8c1eb",
		"fc5c082967dc26",
		"sandbox.smtp.mailtrap.io",
	)

	s.msg = []byte("To: " + email[0] + "\r\n" +
		"Subject: Assunto do email\r\n" +
		"\r\n" +
		"Este é o corpo do email.\r\n")

	err := smtp.SendMail("sandbox.smtp.mailtrap.io:587", auth, "diogosgn@gmail.com", email, s.msg)

	if err != nil {
		log.Printf("Erro ao enviar o email: %v", err)
		return err
	}

	log.Printf("Email enviado com sucesso")

	return nil
}
