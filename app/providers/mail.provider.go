package providers

import (
	"github.com/hirsch88/go-micro-framework/app/mail"
	"github.com/hirsch88/go-micro-framework/config"
	"go.uber.org/zap"
	"log"
)

type mailProvider struct {
	config   *config.MailConfig
	log      *zap.SugaredLogger
	template TemplateProvider
	smtpMail SMTPMailProvider
}

func (p *mailProvider) Send(mail mail.Mailable, to string) bool {
	mailTemplate := mail.Build()
	message, err := p.template.Parse(mailTemplate.TemplatePath, mailTemplate.Context)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.smtpMail.Send(
		SmtpServer{
			Host: p.config.Host,
			Port: p.config.Port,
		},
		Mail{
			From:    p.config.From,
			To:      to,
			Subject: mailTemplate.Subject,
			Body:    message,
		},
		p.config.Password,
	); err != nil {
		return false
	}

	return true
}

type MailProvider interface {
	Send(mail mail.Mailable, to string) bool
}

func NewMailProvider(config *config.MailConfig, log *zap.SugaredLogger, template TemplateProvider, smtpMail SMTPMailProvider) MailProvider {
	return &mailProvider{config, log, template, smtpMail}
}
