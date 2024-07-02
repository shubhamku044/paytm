package routes

import (
	"server/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(groupRouter *gin.RouterGroup) {
	userHandlers := &handler.UserHandler{}
	router := groupRouter.Group("/user")
	{
		router.POST("/signup", userHandlers.SignUp())
		router.POST("/signin", userHandlers.SignIn())
	}
}
