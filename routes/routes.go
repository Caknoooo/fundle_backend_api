package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, EventController controller.EventController, TransaksiController controller.TransaksiController, SeederController controller.SeederController, PenarikanController controller.PenarikanController, jwtService services.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", UserController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		routes.POST("/login", UserController.LoginUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
		routes.POST("/transaksi/:event_id", middleware.Authenticate(jwtService), UserController.CreateTransaksiUser)
		routes.PUT("/transaksi/:event_id", middleware.Authenticate(jwtService), UserController.CreateTransaksiUser)
		routes.GET("/transaksi", middleware.Authenticate(jwtService), UserController.GetTransaksiUser)
	}

	eventRoutes := route.Group("/api/event")
	{
		eventRoutes.POST("", middleware.Authenticate(jwtService), EventController.CreateEvent)
		// eventRoutes.GET("", EventController.GetAllEvent)
		eventRoutes.GET("", EventController.Get3Event)
		eventRoutes.GET("/user/:user_id", middleware.Authenticate(jwtService), EventController.GetAllEventByUserID)
		eventRoutes.GET("/get/:id", EventController.GetEventByID)
		eventRoutes.PUT("/:id", middleware.Authenticate(jwtService), EventController.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.Authenticate(jwtService), EventController.DeleteEvent)
		eventRoutes.GET("/like/:user_id/:event_id", middleware.Authenticate(jwtService), EventController.LikeEventByEventID)
		eventRoutes.GET("/last/:event_id", EventController.GetAllEventLastTransaksi)
	}

	transaksiRoutes := route.Group("/api/transaksi")
	{
		transaksiRoutes.GET("", TransaksiController.GetAllTransaksi)
		transaksiRoutes.GET("/get/:id", TransaksiController.GetTransaksiByID)
	}

	seederRoutes := route.Group("/api/seeder")
	{
		seederRoutes.GET("/category", SeederController.GetAllCategory)
		seederRoutes.GET("/bank", SeederController.GetAllBank)
		seederRoutes.GET("/status_pembayaran", SeederController.GetAllStatusPembayaran)
		seederRoutes.GET("/category/:id", SeederController.GetCategoryByID)
		seederRoutes.GET("/bank/:id", SeederController.GetBankByID)
		seederRoutes.GET("/status_pembayaran/:id", SeederController.GetStatusPembayaranByID)
	}

	penarikanRoutes := route.Group("/api/penarikan")
	{
		penarikanRoutes.POST("", middleware.Authenticate(jwtService), PenarikanController.CreatePenarikan)
		penarikanRoutes.GET("", middleware.Authenticate(jwtService), PenarikanController.GetPenarikanByUser)
	}
}
