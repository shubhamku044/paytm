package routes

import (
	"server/api/controllers"
	"server/api/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controllers.UserController
}

func (r *UserRoutes) InitRoutes(groupRouter *gin.RouterGroup, userController *controllers.UserController) {
	r.userController = userController
	router := groupRouter.Group("/user")
	{
		router.POST("/signup", r.userController.SignUp())
		router.POST("/signin", r.userController.SignIn())
		router.Use(middleware.CheckMiddleware)
		router.GET("/:username", r.userController.GetUser())
		router.GET("/accounts", r.userController.GetUserAccounts())
	}
}
