package repositories

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/app/providers"
)

type userRepository struct {
	database providers.DatabaseProvider
}

func (r *userRepository) Create(user models.User) models.User {
	db := r.database.Connect()
	defer db.Close()

	db.Create(&user)
	return user
}

type UserRepository interface {
	Create(user models.User) models.User
}

func NewUserRepository(database providers.DatabaseProvider) UserRepository {
	return &userRepository{
		database,
	}
}
