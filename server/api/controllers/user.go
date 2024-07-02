package controllers

import (
	"fmt"
	"net/http"
	"server/api/models"
	"server/api/services"
	"server/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserServices
}

func (u *UserController) userRoutes(groupRouter *gin.RouterGroup) {
	router := groupRouter.Group("/user")
	{
		router.POST("/signup", u.SignUp())
		router.POST("/signin", u.SignIn())
	}
}

func (u *UserController) InitUserControllerRoutes(groupRouter *gin.RouterGroup, userServices services.UserServices) {
	u.userRoutes(groupRouter)
	u.userService = &userServices
}

func (u *UserController) SignUp() gin.HandlerFunc {
	type userBody struct {
		FirstName       string `json:"firstName" binding:"required"`
		LastName        string `json:"lastName" binding:"required"`
		UserName        string `json:"userName" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}
	return func(c *gin.Context) {
		var user userBody

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
			return
		}

		valid, errorMsg := utils.ValidateUser(
			user.FirstName,
			user.LastName,
			user.UserName,
			user.Password,
			user.ConfirmPassword,
		)

		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMsg,
			})
			return
		}

		resp, err := u.userService.CreateUserService(&models.User{
			FirstName: strings.TrimSpace(user.FirstName),
			LastName:  strings.TrimSpace(user.LastName),
			UserName:  strings.TrimSpace(user.UserName),
			Password:  strings.TrimSpace(user.Password),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, err := utils.CreateToken(resp.UserName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"token":   token,
		})
	}
}

func (u *UserController) SignIn() gin.HandlerFunc {
	type userBody struct {
		UserName string `json:"userName" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		var user userBody

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
			return
		}

		fmt.Println("Sign in")

		c.JSON(200, gin.H{
			"message": "Sign in",
		})
	}
}
