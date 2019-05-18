package controllers

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/models"
)

type UserController struct {
	db func() *gorm.DB
}

func (c *UserController) Create(ctx *gin.Context) {
	log.Info("STARTING UserController.create()")
	db := c.db()
	defer db.Close()

	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&user)
	ctx.JSON(http.StatusCreated, user)
	log.WithField("username", user.Username).Info("FINISHED UserController.create()")
}

func NewUserController(db func() *gorm.DB) *UserController {
	return &UserController{
		db,
	}
}
