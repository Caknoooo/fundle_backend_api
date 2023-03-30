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
		pembayaranRepository repository.PembayaranRepository = repository.NewPembayaranRepository(db)
		pembayaranService    services.PembayaranService      = services.NewPembayaranService(pembayaranRepository)
		transaksiRepository  repository.TransaksiRepository  = repository.NewTransaksiRepository(db)
		transaksiService     services.TransaksiService       = services.NewTransaksiService(transaksiRepository)
		transaksiController  controller.TransaksiController  = controller.NewTransaksiController(transaksiService, jwtService)
		eventRepository      repository.EventRepository      = repository.NewEventRepository(db)
		eventService         services.EventService           = services.NewEventRepository(eventRepository)
		eventController      controller.EventController      = controller.NewEventController(eventService, transaksiService, jwtService)
		userRepository       repository.UserRepository       = repository.NewUserRepository(db)
		userService          services.UserService            = services.NewUserService(userRepository)
		userController       controller.UserController       = controller.NewUserController(userService, transaksiService, pembayaranService, eventService, jwtService)
	)

	server := gin.Default()
	routes.Router(server, userController, eventController, transaksiController, jwtService)
	server.Use(middleware.CORSMiddleware())

	if err := config.Seeder(db); err != nil {
		log.Fatalf("error seeding database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
