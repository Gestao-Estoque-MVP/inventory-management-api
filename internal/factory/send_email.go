package factory

type ISendEmail interface {
	MultiEmail(templateID string, filter interface{}) (string, error)
	SendOneEmail(userID string, templateID, filter interface{}) (string, error)
	SendEmail(templateID string, filter interface{}) (string, error)
}

func MultiEmail(filter interface{}) (ISendEmail, error) {}

func SendOneEmail(userID string, templateID string, filter interface{}) (string, error) {}
