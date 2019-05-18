package repositories

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db func() *gorm.DB
}

func (r *UserRepository) Create(user models.User) models.User {
	logrus.Info("STARTING UserController.create()")
	db := r.db()
	defer db.Close()

	db.Create(&user)
	return user
}

func NewUserRepository(db func() *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}
