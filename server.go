package main

import (
	"github.com/amirphl/go-gin-gorm-postgres/config"
	"github.com/amirphl/go-gin-gorm-postgres/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDB()
	authController controller.AuthController = controller.CreateAuthController()
)

func main() {
	defer config.CloseDB(db)

	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	// This is just a block!
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
