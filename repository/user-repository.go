package repository

import (
	"github.com/amirphl/go-gin-gorm-postgres/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// UserRepository ...
type UserRepository interface {
	Create(user entity.User) entity.User
	Update(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) interface{}
	FindByID(userID uint64) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

// CreateUserRepo ...
func CreateUserRepo(db *gorm.DB) UserRepository {
	return &userConnection{db}
}

// `user.Password` must be a plain password if set!
func (u *userConnection) Create(user entity.User) entity.User {
	log.Printf("before Create: %v\n", user)
	user.Password = hashAndSalt([]byte(user.Password))
	u.connection.Save(&user)
	return user
}

// `user.Password` must be a plain password if set!
func (u *userConnection) Update(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var t entity.User
		u.connection.Find(&t, user.ID)
		user.Password = t.Password
	}

	u.connection.Save(&user)
	return user
}

func (u *userConnection) VerifyCredential(email string, password string) interface{} {
	res := u.FindByEmail(email)
	user, ok := res.(entity.User)

	comparedPass := comparePassword([]byte(user.Password), []byte(password))

	if comparedPass && ok {
		return user
	}

	log.Printf("user credential verification failed:\t %v \n", email)
	return nil
}

func (u *userConnection) FindByEmail(email string) interface{} {
	var user entity.User

	res := u.connection.Where("email = ?", email).Take(&user) // TODO get, filter

	if res.Error != nil {
		log.Printf("user not found:\t %v \t %v \t %T \t %v \t %T \n", email, res.Error, res.Error, res, res)
		return nil
	}

	return user
}

func (u *userConnection) FindByID(userID uint64) interface{} {
	var user entity.User

	res := u.connection.Find(&user, userID) // TODO get, filter, no id usage?

	if res.Error != nil {
		log.Printf("user not found:\t %v \t %v \t %T \t %v \t %T \n", userID, res.Error, res.Error, res, res)
		return nil
	}

	return user
}

func hashAndSalt(plainPwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(plainPwd, bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func comparePassword(hashed []byte, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, plain)

	if err != nil {
		return false
	}

	return true
}
