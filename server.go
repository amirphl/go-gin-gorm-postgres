package main

import (
	"github.com/amirphl/go-gin-gorm-postgres/config"
	"github.com/amirphl/go-gin-gorm-postgres/controller"
	// "github.com/amirphl/go-gin-gorm-postgres/middleware"
	"github.com/amirphl/go-gin-gorm-postgres/repository"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDB()
	userRepo       repository.UserRepository = repository.CreateUserRepo(db)
	jwtSer         service.JWTService        = service.CreateJWTService()
	authController controller.AuthController = controller.CreateAuthController(userRepo, jwtSer)
)

func main() {
	defer config.CloseDB(db)

	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	// authRoutes := r.Group("api/v1/auth", middleware.AuthorizeJWT(jwtSer))
	// This is just a block!
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
