package main

import (
	"log"
	"os"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/routes"
	"github.com/Caknoooo/golang-clean_template/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {	
	var (
		db             *gorm.DB                  = config.SetUpDatabaseConnection()
		jwtService     services.JWTService       = services.NewJWTService()
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    services.UserService      = services.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)
	)

	if err := config.Seeder(db); err != nil {
		log.Fatalf("error seeding database: %v", err)
	}

	server := gin.Default()
	routes.Router(server, userController, jwtService)
	server.Use(middleware.CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
