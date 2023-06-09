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
		eventService         services.EventService           = services.NewEventService(eventRepository)
		eventController      controller.EventController      = controller.NewEventController(eventService, transaksiService, jwtService, db)
		userRepository       repository.UserRepository       = repository.NewUserRepository(db)
		userService          services.UserService            = services.NewUserService(userRepository)
		userController       controller.UserController       = controller.NewUserController(userService, transaksiService, pembayaranService, eventService, db, jwtService)
		seederRepository     repository.SeederRepository     = repository.NewSeederRepository(db)
		seederService        services.SeederService          = services.NewSeederService(seederRepository)
		seederController     controller.SeederController     = controller.NewSeederController(seederService)
		penarikanRepository  repository.PenarikanRepository  = repository.NewPenarikanRepository(db)
		penarikanService services.PenarikanService = services.NewPenarikanService(penarikanRepository)
		penarikanController controller.PenarikanController = controller.NewPenarikanController(userService, eventService, penarikanService, db, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.Router(server, userController, eventController, transaksiController, seederController, penarikanController, jwtService)

	if err := config.Seeder(db); err != nil {
		log.Fatalf("error seeding database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}