package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIController struct{}

func (c *APIController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func NewAPIController() *APIController {
	return &APIController{}
}
