package config

import (
	"github.com/hirsch88/go-micro-framework/lib"
)

func Mail() *MailConfig {
	return &MailConfig{
		/*
		|--------------------------------------------------------------------------
		| Mail Configurations
		|--------------------------------------------------------------------------
		|
		| In our case we decided to use the GMail SMTP Server. Just add your credentials
		| below or in the .env file.
		|
		*/

		Host:     lib.Env("MAIL_HOST", "smtp.googlemail.com"),
		Port:     lib.Env("MAIL_PORT", "465"),
		From:     lib.Env("MAIL_FROM", "example@gmail.com"),
		Username: lib.Env("MAIL_USERNAME", "username"),
		Password: lib.Env("MAIL_PASSWORD", "password"),
	}
}

type MailConfig struct {
	Host     string
	Port     string
	From     string
	Username string
	Password string
}
