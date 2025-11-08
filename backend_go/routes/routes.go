package routes

import (
	"github.com/gin-gonic/gin"
	bookingcontroller "github.com/rahmat-ama/badminton_reservation/controllers/booking_controller"
	courtcontroller "github.com/rahmat-ama/badminton_reservation/controllers/court_controller"
	timeslotcontroller "github.com/rahmat-ama/badminton_reservation/controllers/timeslot_controller"
	usercontroller "github.com/rahmat-ama/badminton_reservation/controllers/user_controller"
	"github.com/rahmat-ama/badminton_reservation/middleware"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", usercontroller.RegisterUser) // default customer role
			auth.POST("/login", usercontroller.LoginUser)
		}

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			user := protected.Group("/user")
			{
				user.GET("", middleware.AdminOnly(), usercontroller.Get)
				user.GET("/:id", middleware.AdminOrCustomer(), usercontroller.Show)
				user.PUT("/:id", middleware.AdminOrCustomer(), usercontroller.Update)
				user.DELETE("/:id", middleware.AdminOnly(), usercontroller.Delete)
			}

			court := protected.Group("/court")
			{
				court.GET("", courtcontroller.Index)
				court.GET("/:id", courtcontroller.Show)
				court.POST("", middleware.AdminOnly(), courtcontroller.Create)
				court.PUT("/:id", middleware.AdminOnly(), courtcontroller.Update)
				court.DELETE("/:id", middleware.AdminOnly(), courtcontroller.Delete)
			}

			timeslot := protected.Group("/timeslot")
			{
				timeslot.GET("", timeslotcontroller.Index)
				timeslot.GET("/:id", timeslotcontroller.Show)
				timeslot.POST("", middleware.AdminOnly(), timeslotcontroller.Create)
				timeslot.PUT("/:id", middleware.AdminOnly(), timeslotcontroller.Update)
				timeslot.DELETE("/:id", middleware.AdminOnly(), timeslotcontroller.Delete)
			}

			booking := protected.Group("/booking")
			{
				booking.GET("", bookingcontroller.Get)
				booking.GET("/:id", bookingcontroller.Show)
				booking.POST("", middleware.CustomerOnly(), bookingcontroller.Create)
				booking.PUT("/:id", middleware.AdminOrCustomer(), bookingcontroller.Update)
				booking.DELETE("/:id", middleware.AdminOnly(), bookingcontroller.Delete)
			}
		}
	}
}
