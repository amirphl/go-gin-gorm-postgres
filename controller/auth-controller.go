package controller

import (
	"github.com/amirphl/go-gin-gorm-postgres/dto"
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"github.com/amirphl/go-gin-gorm-postgres/helper"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
)

// AuthController ...
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authSer service.AuthService
	jwtSer  service.JWTService
}

func (a *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO

	if err := ctx.ShouldBind(&loginDTO); err != nil {
		resp := helper.BuildErrResp("Failed to process login request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if err := validator.Validate(loginDTO); err != nil {
		resp := helper.BuildErrResp("Failed to validate login request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	res := a.authSer.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if user, ok := res.(entity.User); ok {
		genToken := a.jwtSer.GenerateToken(user.ID)
		user.Token = genToken
		resp := helper.BuildResp("Login OK!", user)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp := helper.BuildErrResp("Invalid login credential", "", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
}

func (a *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	if err := ctx.ShouldBind(&registerDTO); err != nil {
		resp := helper.BuildErrResp("Failed to process register request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if err := validator.Validate(registerDTO); err != nil {
		resp := helper.BuildErrResp("Failed to validate register request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	res := a.authSer.FindUserByEmail(registerDTO.Email)

	if _, ok := res.(entity.User); ok {
		resp := helper.BuildErrResp("Duplicate email", "", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, resp)
		return
	}

	newUser := a.authSer.CreateUser(registerDTO)
	genToken := a.jwtSer.GenerateToken(newUser.ID)
	newUser.Token = genToken
	resp := helper.BuildResp("register OK!", newUser)
	ctx.JSON(http.StatusCreated, resp)
}

// CreateAuthController ...
func CreateAuthController(authSer service.AuthService, jwtSer service.JWTService) AuthController {
	return &authController{
		authSer,
		jwtSer,
	}
}
