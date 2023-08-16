package factory

import (
	"fmt"
)

type ISendEmail interface {
	MultiEmail() (string, error)
	SendOneEmail() (string, error)
}

type SendEmailFactory struct{}

const (
	internal = "internal"
	external = "external"
)

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
