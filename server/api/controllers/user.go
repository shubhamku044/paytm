package controllers

import (
	"fmt"
	"server/api/models"
	"server/api/routes"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) InitUserControllerRoutes(groupRouter *gin.RouterGroup) {
	routes.UserRoutes(groupRouter)
}

func (u *UserController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)

		if err != nil {
			fmt.Println("Error binding JSON")
		}

		c.JSON(200, gin.H{
			"message": "Sign up",
		})
	}
}

func (u *UserController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Sign in")

		c.JSON(200, gin.H{
			"message": "Sign in",
		})

	}
}
