package main

import (
	"fmt"
	"net/http"
	"os"
	"server/api/controllers"
	"server/api/routes"
	"server/api/services"
	"server/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		os.Exit(1)
	}

	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	db := database.ConnectDB()

	v1 := r.Group("/api/v1")

	// services
	userServices := &services.UserServices{}
	userServices.InitServices(db)

	accountServices := &services.AccountServices{}
	accountServices.InitServices(db)

	// controllers
	userController := &controllers.UserController{}
	userController.InitController(userServices)

	accountController := &controllers.AccountController{}
	accountController.InitController(accountServices)

	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})

		userRoutes := &routes.UserRoutes{}
		userRoutes.InitRoutes(v1, userController)

		accountRoutes := &routes.AccountRoutes{}
		accountRoutes.InitRoutes(v1, accountController)
	}

	r.Run(":8080")
}
