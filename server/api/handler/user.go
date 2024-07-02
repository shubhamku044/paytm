package handler

import (
	"fmt"
	"server/api/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (u *UserHandler) SignUp() gin.HandlerFunc {
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

func (u *UserHandler) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Sign in")

		c.JSON(200, gin.H{
			"message": "Sign in",
		})

	}
}
