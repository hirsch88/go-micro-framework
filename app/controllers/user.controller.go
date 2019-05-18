package controllers

import (
	"github.com/hirsch88/go-micro-framework/lib"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct {}

func (c *UserController) Create(ctx *gin.Context) {
	db := lib.DB()
	defer db.Close()

	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&user)
	ctx.JSON(http.StatusCreated, user)
}

func NewUserController() *UserController {
	return &UserController{}
}
