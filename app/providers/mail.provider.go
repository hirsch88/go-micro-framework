package providers

import (
	"bytes"
	"github.com/hirsch88/go-micro-framework/app/mail"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/core"
	"html/template"
	"log"
)

type mailProvider struct {
	config *config.MailConfig
}

func (p *mailProvider) Send(mail mail.Mailable, to string) {
	mailTemplate := mail.Build()
	message, err := p.parseTemplate(mailTemplate)
	if err != nil {
		log.Fatal(err)
	}
	if ok := p.sendMail(to, mailTemplate.Subject, message); ok {
		log.Printf("Email has been sent to %s\n", to)
	} else {
		log.Printf("Failed to send the email to %s\n", to)
	}
}

func (p *mailProvider) parseTemplate(mailTemplate *mail.Template) (string, error) {
	t, err := template.ParseFiles(mailTemplate.TemplatePath)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, mailTemplate.Context); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (p *mailProvider) sendMail(to string, subject string, message string) bool {
	smtpServer := core.SmtpServer{
		Host: p.config.Host,
		Port: p.config.Port,
	}

	mail := core.Mail{
		From:    p.config.From,
		To:      to,
		Subject: subject,
		Body:    message,
	}

	if err := core.SendMail(smtpServer, mail, p.config.Password); err != nil {
		return false
	}

	return true
}

type MailProvider interface {
	Send(mail mail.Mailable, to string)
}

func NewMailProvider(config *config.MailConfig) MailProvider {
	return &mailProvider{config}
}
