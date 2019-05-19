package providers

import (
	"github.com/hirsch88/go-micro-framework/lib"
)

type mailProvider struct {
	host     string
	port     string
	from     string
	password string
}

func (m *mailProvider) SendMail(to string, subject string, body string) {

	smtpServer := lib.SmtpServer{
		Host: m.host,
		Port: m.port,
	}

	mail := lib.Mail{
		From:    m.from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	lib.SendMail(smtpServer, mail, m.password)

}

type MailProvider interface {
	SendMail(to string, subject string, body string)
}

func NewMailProvider(host string, port string, from string, password string) MailProvider {
	return &mailProvider{host, port, from, password}
}
