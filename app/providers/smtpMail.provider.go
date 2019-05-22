package providers

import (
	"crypto/tls"
	"fmt"
	"go.uber.org/zap"
	"net/smtp"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type smtpMailProvider struct {
	log *zap.SugaredLogger
}

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.From)
	message += fmt.Sprintf("To: %s\r\n", mail.To)
	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += fmt.Sprintf("%s\r\n", MIME)
	message += mail.Body
	return message
}

func (p *smtpMailProvider) Send(smtpServer SmtpServer, mail Mail, password string) error {
	messageBody := mail.BuildMessage()
	auth := smtp.PlainAuth("", mail.From, password, smtpServer.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	// Gmail will reject connection if it's not secure
	// TLS config
	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.From); err != nil {
		return err
	}
	if err = client.Rcpt(mail.To); err != nil {
		return err
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()
	return nil
}

type SMTPMailProvider interface {
	Send(smtpServer SmtpServer, mail Mail, password string) error
}

func NewSMTPMailProvider(log *zap.SugaredLogger) SMTPMailProvider {
	return &smtpMailProvider{log}
}
