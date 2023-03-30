package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, EventController controller.EventController, TransaksiController controller.TransaksiController, jwtService services.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", UserController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		routes.POST("/login", UserController.LoginUser)
		routes.POST("/logout", middleware.Authenticate(jwtService), UserController.LogoutUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
		routes.POST("/event/:event_id", middleware.Authenticate(jwtService), UserController.CreateTransaksiUser)
		routes.PUT("/event/:event_id", middleware.Authenticate(jwtService), UserController.CreateTransaksiUser)
		routes.GET("/transaksi", middleware.Authenticate(jwtService), UserController.GetTransaksiUser)
	}

	eventRoutes := route.Group("/api/event")
	{
		eventRoutes.POST("", middleware.Authenticate(jwtService), EventController.CreateEvent)
		eventRoutes.GET("", EventController.GetAllEvent)                                                            // Permission -> Admin?
		eventRoutes.GET("/user/:user_id", middleware.Authenticate(jwtService), EventController.GetAllEventByUserID) // Hanya user yang membuatnya yang bisa cek seluruh event yang dia buat
		eventRoutes.GET("/get/:id", EventController.GetEventByID)
		eventRoutes.PUT("/:id", middleware.Authenticate(jwtService), EventController.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.Authenticate(jwtService), EventController.DeleteEvent)
		eventRoutes.GET("/like/:user_id/:event_id", middleware.Authenticate(jwtService), EventController.LikeEventByEventID)
		eventRoutes.GET("/:event_id", EventController.GetAllEventLastTransaksi)
	}

	transaksiRoutes := route.Group("/api/transaksi")
	{
		transaksiRoutes.GET("", TransaksiController.GetAllTransaksi) // Permission -> Admin?
		transaksiRoutes.GET("/get/:id", TransaksiController.GetTransaksiByID)
		transaksiRoutes.GET("/:user_id", middleware.Authenticate(jwtService), TransaksiController.GetTransaksiByUserID)
	}
}
