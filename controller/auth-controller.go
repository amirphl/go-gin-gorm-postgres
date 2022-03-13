package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth Controller is a contract
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func (a *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok login",
	})
}

func (a *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok register",
	})
}

// TODO pointer type
func CreateAuthController() AuthController {
	return &authController{}
}
