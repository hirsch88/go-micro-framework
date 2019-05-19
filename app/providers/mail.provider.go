package providers

import (
	"bytes"
	"github.com/hirsch88/go-micro-framework/lib"
	"html/template"
	"log"
)

type mailProvider struct {
	host     string
	port     string
	from     string
	password string
}

func (p *mailProvider) Send(mailTemplate string, to string, subject string, data interface{}) {
	message, err := p.parseTemplate(mailTemplate, data)
	if err != nil {
		log.Fatal(err)
	}
	if ok := p.sendMail(to, subject, message); ok {
		log.Printf("Email has been sent to %s\n", to)
	} else {
		log.Printf("Failed to send the email to %s\n", to)
	}
}

func (p *mailProvider) parseTemplate(mailTemplate string, data interface{}) (string, error) {
	t, err := template.ParseFiles(mailTemplate)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (p *mailProvider) sendMail(to string, subject string, message string) bool {
	smtpServer := lib.SmtpServer{
		Host: p.host,
		Port: p.port,
	}

	mail := lib.Mail{
		From:    p.from,
		To:      to,
		Subject: subject,
		Body:    message,
	}

	if err := lib.SendMail(smtpServer, mail, p.password); err != nil {
		return false
	}

	return true
}

type MailProvider interface {
	Send(mailTemplate string, to string, subject string, data interface{})
}

func NewMailProvider(host string, port string, from string, password string) MailProvider {
	return &mailProvider{host, port, from, password}
}
