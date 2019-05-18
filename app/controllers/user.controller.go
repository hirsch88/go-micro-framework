package controllers

import (
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct {
	db func() *gorm.DB
}

func (c *UserController) Create(ctx *gin.Context) {
	db := c.db()
	defer db.Close()

	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&user)
	ctx.JSON(http.StatusCreated, user)
}

func NewUserController(db func() *gorm.DB) *UserController {
	return &UserController{
		db,
	}
}
