package factory

import "fmt"

type SendEmailExternal struct{}

func (si *SendEmailExternal) MultiEmail() (string, error) {
	return "", fmt.Errorf("Error")

}

func (si *SendEmailExternal) SendOneEmail() (string, error) {
	return "", fmt.Errorf("Error")

}
