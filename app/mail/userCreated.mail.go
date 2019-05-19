package mail

import "github.com/hirsch88/go-micro-framework/app/models"

type userCreatedMail struct {
	user models.User
}

func (m *userCreatedMail) Build() *Template {
	return &Template{
		TemplatePath: "resources/mail/userCreated.mail.html",
		Subject:      "Registration Confirmation",
		Context:      map[string]string{"username": m.user.Username},
	}
}

func NewUserCreatedMail(user models.User) Mailable {
	return &userCreatedMail{user}
}
