package config

func NewMailConfig() *MailConfig {
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

		Host:     Env("MAIL_HOST", "smtp.googlemail.com"),
		Port:     Env("MAIL_PORT", "465"),
		From:     Env("MAIL_FROM", "example@gmail.com"),
		Username: Env("MAIL_USERNAME", "username"),
		Password: Env("MAIL_PASSWORD", "password"),
	}
}

type MailConfig struct {
	Host     string
	Port     string
	From     string
	Username string
	Password string
}
