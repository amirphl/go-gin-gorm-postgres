package service

import (
	"github.com/amirphl/go-gin-gorm-postgres/dto"
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"github.com/amirphl/go-gin-gorm-postgres/repository"
	"github.com/mashingan/smapping"
	"log"
)

// AuthService ... UserRepository
type AuthService interface {
	CreateUser(user dto.RegisterDTO) entity.User
	UpdateUser(user dto.UserUpdateDTO) entity.User
	VerifyCredential(email string, password string) interface{}
	FindUserByEmail(email string) interface{}
	FindUserByID(userID uint64) interface{}
}

type authService struct {
	userRepo repository.UserRepository
}

// CreateAuthService ...
func CreateAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (a *authService) CreateUser(regDTO dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&regDTO))

	if err != nil {
		log.Fatalf("failed to map RegisterDTO to User: \t %v \t %v \t %T", regDTO, err, err)
		// TODO
	}

	return a.userRepo.Insert(userToCreate)
}

func (a *authService) UpdateUser(user dto.UserUpdateDTO) entity.User {
	return entity.User{} // TODO NotImplemented
}

func (a *authService) VerifyCredential(email string, password string) interface{} {
	return a.userRepo.VerifyCredential(email, password)
}

func (a *authService) FindUserByEmail(email string) interface{} {
	return a.userRepo.FindByEmail(email)
}

func (a *authService) FindUserByID(userID uint64) interface{} {
	return a.userRepo.FindByID(userID)
}
