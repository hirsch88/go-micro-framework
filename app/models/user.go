package models

type User struct {
	Identifiable
	UseTimestamp

	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
