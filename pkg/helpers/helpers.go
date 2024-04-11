package helpers

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func OnlyName(name string) string {
	return strings.Trim(name, " ")
}

func OnlyDocument(document string) string {
	regexp := regexp.MustCompile("[^0-9]+")
	return regexp.ReplaceAllString(document, "")
}

func OnlyEmail(email string) string {
	return strings.Trim(email, " ")
}

func OnlyMobilePhone(mobilePhone string) string {
	regexp := regexp.MustCompile("[^0-9]+")
	return regexp.ReplaceAllString(mobilePhone, "")
}

func OnlyPassword(password string) string {
	return strings.Trim(password, " ")
}

func OnlyNumbers(numbers string) string {
	regexp := regexp.MustCompile("[^0-9]+")
	return regexp.ReplaceAllString(numbers, "")
}

func VerifyPassword(password, verify string) bool {
	hash := bcrypt.CompareHashAndPassword([]byte(verify), []byte(password))

	return hash == nil
}
