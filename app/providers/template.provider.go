package providers

import (
	"bytes"
	"go.uber.org/zap"
	"html/template"
)

type templateProvider struct {
	log *zap.SugaredLogger
}

func (p *templateProvider) Parse(templatePath string, context interface{}) (string, error) {
	p.log.Infof("Parsing template %s", templatePath)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, context); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

type TemplateProvider interface {
	Parse(templatePath string, context interface{}) (string, error)
}

func NewTemplateProvider(log *zap.SugaredLogger) TemplateProvider {
	return &templateProvider{log}
}
