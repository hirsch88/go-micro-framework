package services

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/app/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

func (s *userService) Create(user models.User) models.User {
	return s.userRepository.Create(user)
}

type UserService interface {
	Create(user models.User) models.User
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository,
	}
}
