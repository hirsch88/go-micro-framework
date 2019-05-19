package lib

type Mailable interface {
	Build() *MailTemplate
}

type MailTemplate struct {
	TemplatePath string
	Subject      string
	Context      interface{}
}
