package mail

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/lib"
)

type userCreatedMail struct {
	user models.User
}

func (m *userCreatedMail) Build() *lib.MailTemplate {
	return &lib.MailTemplate{
		TemplatePath: "resources/mail/userCreated.mail.html",
		Subject:      "Registration Confirmation",
		Context:      map[string]string{"username": m.user.Username},
	}
}

func NewUserCreatedMail(user models.User) lib.Mailable {
	return &userCreatedMail{user}
}
