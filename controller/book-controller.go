package controller

import (
	"net/http"
	"strconv"

	"github.com/amirphl/go-gin-gorm-postgres/dto"
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"github.com/amirphl/go-gin-gorm-postgres/helper"
	"github.com/amirphl/go-gin-gorm-postgres/repository"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"gopkg.in/validator.v2"
)

// BookController ...
type BookController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Retrive(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	userRepo repository.UserRepository
	bookRepo repository.BookRepository
	jwtSer   service.JWTService
}

func (b *bookController) Create(ctx *gin.Context) {
	var dto dto.BookCreateDTO
	err1 := ctx.ShouldBind(&dto)
	err2 := validator.Validate(dto)

	if err1 != nil || err2 != nil {
		resp := helper.BuildErrResp(helper.DefaultErrMsg, "", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	user, ok := b.jwtSer.ExtractUser(authHeader, b.userRepo).(entity.User)

	if !ok {
		resp := helper.BuildErrResp("User does not exist", "", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	dto.UserID = user.ID
	bookToCreate := entity.Book{}

	if err := smapping.FillStruct(&bookToCreate, smapping.MapFields(&dto)); err != nil {
		resp := helper.BuildErrResp("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	newBook := b.bookRepo.Create(bookToCreate)
	resp := helper.BuildResp("Book created!", newBook)
	ctx.JSON(http.StatusCreated, resp)
}

func (b *bookController) Update(ctx *gin.Context) {
	var dto dto.BookUpdateDTO
	err1 := ctx.ShouldBind(&dto)
	err2 := validator.Validate(dto)
	bookID, err3 := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		resp := helper.BuildErrResp(helper.DefaultErrMsg, "", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	user, userExists := b.jwtSer.ExtractUser(authHeader, b.userRepo).(entity.User)

	if !userExists {
		resp := helper.BuildErrResp("User does not exist", "", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	isOwner := b.bookRepo.IsOwner(user.ID, bookID)

	if !isOwner {
		resp := helper.BuildErrResp("Action is forbidden", "", nil)
		ctx.AbortWithStatusJSON(http.StatusForbidden, resp)
		return
	}

	dto.UserID = user.ID
	dto.ID = bookID
	bookToUpdate := entity.Book{}

	if err := smapping.FillStruct(&bookToUpdate, smapping.MapFields(&dto)); err != nil {
		resp := helper.BuildErrResp("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	updatedBook := b.bookRepo.Update(bookToUpdate)
	resp := helper.BuildResp("Book updated!", updatedBook)
	ctx.JSON(http.StatusNoContent, resp)
}

func (b *bookController) Retrive(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		resp := helper.BuildErrResp(helper.DefaultErrMsg, "", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	book, ok := b.bookRepo.Retrive(bookID).(entity.Book)

	if !ok {
		resp := helper.BuildErrResp("Book not found", "", nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
		return
	}

	resp := helper.BuildResp("Book found!", book)
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

// TODO pagination
func (b *bookController) List(ctx *gin.Context) {
	books := b.bookRepo.List()
	resp := helper.BuildResp("Books found!", books)
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func (b *bookController) Delete(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		resp := helper.BuildErrResp(helper.DefaultErrMsg, "", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	book, ok := b.bookRepo.Retrive(bookID).(entity.Book)

	if !ok {
		resp := helper.BuildErrResp("Book not found", "", nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	user, userExists := b.jwtSer.ExtractUser(authHeader, b.userRepo).(entity.User)

	if !userExists {
		resp := helper.BuildErrResp("User does not exist", "", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	isOwner := b.bookRepo.IsOwner(user.ID, bookID)

	if !isOwner {
		resp := helper.BuildErrResp("Action is forbidden", "", nil)
		ctx.AbortWithStatusJSON(http.StatusForbidden, resp)
		return
	}

	b.bookRepo.Delete(book)
	resp := helper.BuildResp("Book Deleted!", book)
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

// CreateBookController ...
func CreateBookController(userRepo repository.UserRepository, bookRepo repository.BookRepository, jwtSer service.JWTService) BookController {
	return &bookController{userRepo, bookRepo, jwtSer}
}
