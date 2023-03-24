package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, EventController controller.EventController, jwtService services.JWTService) {
func Router(route *gin.Engine, UserController controller.UserController, EventController controller.EventController, jwtService services.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", UserController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		routes.POST("/login", UserController.LoginUser)
		routes.POST("/logout", middleware.Authenticate(jwtService), UserController.LogoutUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
		routes.GET("/event", middleware.Authenticate(jwtService), UserController.CreateTransaksiUser)
	}
}

func RouterTransaksi(route *gin.Engine, TransaksiController controller.TransaksiController) {
	routes := route.Group("/api/transaksi")
	{
		routes.GET("/event", TransaksiController.GetAllTransaksi)
	}

	eventRoutes := route.Group("/api/event")
	{
		eventRoutes.POST("", middleware.Authenticate(jwtService), EventController.CreateEvent)
		eventRoutes.GET("", EventController.GetAllEvent) // Permission -> Admin
		eventRoutes.GET("/:user_id", middleware.Authenticate(jwtService), EventController.GetAllEventByUserID) // Hanya user yang membuatnya yang bisa cek seluruh event yang dia buat
		eventRoutes.GET("/get/:id", EventController.GetEventByID)
		eventRoutes.PUT("/:id", middleware.Authenticate(jwtService), EventController.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.Authenticate(jwtService), EventController.DeleteEvent)
		eventRoutes.GET("/like/:user_id/:event_id", middleware.Authenticate(jwtService), EventController.LikeEventByEventID)
	}

	eventRoutes := route.Group("/api/event")
	{
		eventRoutes.POST("", middleware.Authenticate(jwtService), EventController.CreateEvent)
		eventRoutes.GET("", EventController.GetAllEvent) // Permission -> Admin
		eventRoutes.GET("/:user_id", middleware.Authenticate(jwtService), EventController.GetAllEventByUserID) // Hanya user yang membuatnya yang bisa cek seluruh event yang dia buat
		eventRoutes.GET("/get/:id", EventController.GetEventByID)
		eventRoutes.PUT("/:id", middleware.Authenticate(jwtService), EventController.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.Authenticate(jwtService), EventController.DeleteEvent)
		eventRoutes.POST("/like/:user_id/:event_id", middleware.Authenticate(jwtService), EventController.LikeEventByEventID)
	}
}

