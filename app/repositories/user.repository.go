package repositories

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db func() *gorm.DB
}

func (r *userRepository) Create(user models.User) models.User {
	logrus.Info("STARTING UserController.create()")
	db := r.db()
	defer db.Close()

	db.Create(&user)
	return user
}

type UserRepository interface {
	Create(user models.User) models.User
}

func NewUserRepository(db func() *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}
