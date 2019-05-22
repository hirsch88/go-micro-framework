package controllers

import (
	"github.com/hirsch88/go-micro-framework/app/services"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct {
	userService services.UserService
	log         *zap.SugaredLogger
}

func (c *UserController) Create(ctx *gin.Context) {
	c.log.Info("STARTING UserController.create()")

	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = c.userService.Create(user)

	ctx.JSON(http.StatusCreated, user)
	c.log.Infow("FINISHED UserController.create()",
		"username", user.Username,
	)
}

func NewUserController(userService services.UserService, log *zap.SugaredLogger) *UserController {
	return &UserController{
		userService,
		log,
	}
}
