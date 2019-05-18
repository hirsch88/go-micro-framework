package services

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func (s *UserService) Create(user models.User) models.User {
	return s.userRepository.Create(user)
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}
