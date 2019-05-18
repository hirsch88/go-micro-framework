package controllers

import (
	"github.com/hirsch88/go-micro-framework/app/services"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct {
	userService *services.UserService
}

func (c *UserController) Create(ctx *gin.Context) {
	log.Info("STARTING UserController.create()")

	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = c.userService.Create(user)

	ctx.JSON(http.StatusCreated, user)
	log.WithField("username", user.Username).Info("FINISHED UserController.create()")
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService,
	}
}
