package main

import (
	"github.com/amirphl/go-gin-gorm-postgres/config"
	"github.com/amirphl/go-gin-gorm-postgres/controller"
	"github.com/amirphl/go-gin-gorm-postgres/middleware"
	"github.com/amirphl/go-gin-gorm-postgres/repository"
	"github.com/amirphl/go-gin-gorm-postgres/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDB()
	userRepo       repository.UserRepository = repository.CreateUserRepo(db)
	bookRepo       repository.BookRepository = repository.CreateBookRepo(db)
	jwtSer         service.JWTService        = service.CreateJWTService()
	authController controller.AuthController = controller.CreateAuthController(userRepo, jwtSer)
	userController controller.UserController = controller.CreateUserController(userRepo, jwtSer)
	bookController controller.BookController = controller.CreateBookController(userRepo, bookRepo, jwtSer)
)

func main() {
	defer config.CloseDB(db)

	r := gin.Default()

	allRoutes := r.Group("api/v1")
	{
		authRoutes := allRoutes.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}
		protectedUserRoutes := allRoutes.Group("/users", middleware.AuthorizeJWT(jwtSer))
		{
			protectedUserRoutes.PUT("/:id", userController.UpdateUser)
			protectedUserRoutes.GET("/:id", userController.GetUser)
		}
		bookRoutes := allRoutes.Group("/books")
		{
			bookRoutes.GET("/", bookController.List)
			bookRoutes.GET("/:id", bookController.Retrive)
		}
		protectedBookRoutes := allRoutes.Group("/books", middleware.AuthorizeJWT(jwtSer))
		{
			protectedBookRoutes.POST("/", bookController.Create)
			protectedBookRoutes.PUT("/:id", bookController.Update)
			protectedBookRoutes.DELETE("/:id", bookController.Delete)
		}
	}

	r.Run()
}
