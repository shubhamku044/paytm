package controllers

import (
	"net/http"
	"server/api/middleware"
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
		router.Use(middleware.CheckMiddleware)
		router.GET("/:username", u.GetUser())
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

		userExists, err := u.userService.GetUserByUsername(user.UserName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if userExists != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User already exists",
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

		if len(strings.TrimSpace(user.UserName)) < 1 || len(strings.TrimSpace(user.Password)) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Username and password must be provided",
			})
			return
		}

		userDetails, err := u.userService.GetUserByUsername(user.UserName)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		if err := utils.ComparePassword(user.Password, userDetails.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid username or password",
			})
			return
		}

		token, err := utils.CreateToken(userDetails.UserName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User signed in successfully",
			"token":   token,
		})
	}
}

func (u *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		user, err := u.userService.GetUserByUsername(username)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		userNameAfterMiddleware := c.GetString("username")

		if userNameAfterMiddleware != user.UserName {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "You are not authorized to view this user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User found",
			"user":    user,
		})
	}
}
