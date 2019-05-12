package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct{}

func (c *UserController) Create(ctx *gin.Context) {
	var json models.User
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, json)
}

func NewUserController() *UserController {
	return &UserController{}
}
