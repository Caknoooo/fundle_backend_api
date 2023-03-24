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
		db                   *gorm.DB                        = config.SetUpDatabaseConnection()
		jwtService           services.JWTService             = services.NewJWTService()
		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
		transaksiService    services.TransaksiService      = services.NewTransaksiService(transaksiRepository)
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)
		userRepository       repository.UserRepository       = repository.NewUserRepository(db)
		userService          services.UserService            = services.NewUserService(userRepository)
		userController       controller.UserController       = controller.NewUserController(userService, transaksiService, jwtService)
		eventRepository repository.EventRepository = repository.NewEventRepository(db)
		eventService    services.EventService      = services.NewEventRepository(eventRepository)
		eventController controller.EventController = controller.NewEventController(eventService, jwtService)
	)

	server := gin.Default()
	routes.RouterUser(server, userController, eventController, jwtService)
	routes.RouterTransaksi(server, transaksiController)
	server.Use(middleware.CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)

	if err := config.Seeder(db); err != nil {
		log.Fatalf("error seeding database: %v", err)
	}

}
