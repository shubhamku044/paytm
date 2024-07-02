package main

import (
	"fmt"
	"net/http"
	"os"
	"server/api/controllers"
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
	userServices.InitUserServices(db)

	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})

		userController := &controllers.UserController{}
		userController.InitUserControllerRoutes(v1, *userServices)
	}

	r.Run(":8080")
}
