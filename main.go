package main

import (
	"farmconnect/config"
	"farmconnect/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitDB()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://farmconnect-dusky.vercel.app","http://localhost:3001","https://farmconnect.vercel.app/"}, // Change to your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Routes

	//auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/users",controllers.GetUsers)
	r.PUT("/users/:id",controllers.UpdateUser)
	r.DELETE("users/:id",controllers.DeleteUser)
	r.GET("/user/:id",controllers.GetUserByID)

	// Farmer Routes
	r.POST("/farmers", controllers.CreateFarmer)
	r.GET("/farmers", controllers.GetFarmers)
	r.GET("/farmers/:id", controllers.GetFarmerByID)
	r.PUT("/farmers/:id", controllers.UpdateFarmer)
	r.DELETE("/farmers/:id", controllers.DeleteFarmer)

	// Business Routes
	r.POST("/businesses", controllers.CreateBusiness)
	r.GET("/businesses", controllers.GetBusinesses)
	r.GET("/businesses/:id", controllers.GetBusinessByID)
	r.PUT("/businesses/:id", controllers.UpdateBusiness)
	r.DELETE("/businesses/:id", controllers.DeleteBusiness)

	// Logistics Routes
	r.POST("/logistics", controllers.CreateLogistics)
	r.GET("/logistics", controllers.GetLogisticsProviders)
	r.GET("/logistics/:id", controllers.GetLogisticsByID)
	r.PUT("/logistics/:id", controllers.UpdateLogistics)
	r.DELETE("/logistics/:id", controllers.DeleteLogistics)

	// Produce Routes
	r.POST("/produce", controllers.CreateProduce)
	r.GET("/produce", controllers.GetAllProduce)
	r.GET("/produce/:id", controllers.GetProduceByID)
	r.GET("/produce/user/:user_id", controllers.GetProduceByUser)
	r.PUT("/produce/:id", controllers.UpdateProduce)
	r.DELETE("/produce/:id", controllers.DeleteProduce)



	// routes.go (or wherever you define routes)
	r.POST("/api/checkout", controllers.Checkout)          // Create order (Checkout)
	r.GET("/api/orders", controllers.GetAllOrders)         // Get all orders
	r.GET("/api/orders/:id", controllers.GetOrderByID)    // Get order by ID
	r.GET("/api/orders/users/:userID", controllers.GetUserOrders) 
	r.PUT("/api/orders/:id", controllers.UpdateOrderStatus) // Update order status
	r.DELETE("/api/orders/:id", controllers.DeleteOrder)   // Delete order
	r.GET("/orders/user/:user_id", controllers.GetOrdersByUser)


	r.POST("/notifications", controllers.CreateNotification)
	r.GET("/notifications", controllers.GetAllNotifications)
	r.GET("/notifications/user/:user_id", controllers.GetNotificationsByUser)
	r.PUT("/notifications/:id/read", controllers.MarkNotificationAsRead)
	r.DELETE("/notifications/:id", controllers.DeleteNotification)
	


	r.Run(":8000")
}
