package mail

type Mailable interface {
	Build() *Template
}

type Template struct {
	TemplatePath string
	Subject      string
	Context      interface{}
}
