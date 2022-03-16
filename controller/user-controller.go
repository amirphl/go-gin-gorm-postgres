package controller

import (
	"log"
	"net/http"

	"github.com/amirphl/go-gin-gorm-postgres/dto"
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"github.com/amirphl/go-gin-gorm-postgres/helper"
	"github.com/amirphl/go-gin-gorm-postgres/repository"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mashingan/smapping"
	"gopkg.in/validator.v2"
)

// UserController ...
type UserController interface {
	UpdateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}

type userController struct {
	userRepo repository.UserRepository
	jwtSer   service.JWTService
}

// Put this handler behind JWT middleware.
func (u *userController) UpdateUser(ctx *gin.Context) {
	var uuDTO dto.UserUpdateDTO

	if err := ctx.ShouldBind(&uuDTO); err != nil {
		resp := helper.BuildErrResp("Failed to process request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if err := validator.Validate(uuDTO); err != nil {
		resp := helper.BuildErrResp("Failed to process request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	// There is JWT auth middleware between the controller and user request, so no need to check for error.
	token, _ := u.jwtSer.ValidateToken(authHeader)
	userID := uint64(token.Claims.(jwt.MapClaims)["user_id"].(float64))

	if _, ok := u.userRepo.FindByID(userID).(entity.User); !ok {
		resp := helper.BuildErrResp("User does not exist", "", nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
		return
	}

	userToUpdate := entity.User{}

	if err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&uuDTO)); err != nil {
		log.Printf("failed to map UserUpdateDTO to User: \t %v \t %v \t %T", uuDTO, err, err)
		resp := helper.BuildErrResp("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	userToUpdate.ID = userID
	updatedUser := u.userRepo.Update(userToUpdate)
	resp := helper.BuildResp("Update OK!", updatedUser)
	ctx.JSON(http.StatusNoContent, resp) // TODO change to 200?
}

// Put this handler behind JWT middleware.
func (u *userController) GetUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	// There is JWT auth middleware between the controller and user request, so no need to check for error.
	token, _ := u.jwtSer.ValidateToken(authHeader)
	userID := uint64(token.Claims.(jwt.MapClaims)["user_id"].(float64))
	user, ok := u.userRepo.FindByID(userID).(entity.User)

	if !ok {
		resp := helper.BuildErrResp("User does not exist", "", nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
		return
	}

	resp := helper.BuildResp("Get OK!", user)
	ctx.JSON(http.StatusOK, resp)
}

// CreateUserController ...
func CreateUserController(userRepo repository.UserRepository, jwtSer service.JWTService) UserController {
	return &userController{
		userRepo,
		jwtSer,
	}
}
