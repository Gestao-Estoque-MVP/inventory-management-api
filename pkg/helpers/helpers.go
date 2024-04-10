package helpers

import (
	"regexp"
	"strings"
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
