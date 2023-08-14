package factory

type ISendEmail interface {
	MultiEmail(templateID string, filter interface{}) (string, error)
	SendOneEmail(userID string, templateID, filter interface{}) (string, error)
	SendEmail(templateID string, filter interface{}) (string, error)
}

type SendEmailInternal struct {
	to             string
	from           string
	typeTemplate   string
	templateStruct interface{}
}

func (si *SendEmailInternal) MultiEmail() (string, error) {

}

func (si *SendEmailInternal) SendOneEmail(userID string, templateID string, filter interface{}) (string, error) {
}
