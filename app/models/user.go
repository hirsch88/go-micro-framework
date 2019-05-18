package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
