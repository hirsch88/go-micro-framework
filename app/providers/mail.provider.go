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

func (p *mailProvider) Send(to string, subject string, body string) {

	smtpServer := lib.SmtpServer{
		Host: p.host,
		Port: p.port,
	}

	mail := lib.Mail{
		From:    p.from,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	lib.SendMail(smtpServer, mail, p.password)

}

type MailProvider interface {
	Send(to string, subject string, body string)
}

func NewMailProvider(host string, port string, from string, password string) MailProvider {
	return &mailProvider{host, port, from, password}
}
