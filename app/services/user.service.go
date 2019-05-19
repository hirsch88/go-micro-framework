package services

import (
	"github.com/hirsch88/go-micro-framework/app/mail"
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/app/providers"
	"github.com/hirsch88/go-micro-framework/app/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
	mailProvider   providers.MailProvider
}

func (s *userService) Create(user models.User) models.User {
	newUser := s.userRepository.Create(user)
	s.mailProvider.Send(mail.NewUserCreatedMail(user), user.Email)
	return newUser
}

type UserService interface {
	Create(user models.User) models.User
}

func NewUserService(userRepository repositories.UserRepository, mailProvider providers.MailProvider) UserService {
	return &userService{
		userRepository,
		mailProvider,
	}
}
