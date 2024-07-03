package routes

import (
	"server/api/controllers"
	"server/api/middleware"

	"github.com/gin-gonic/gin"
)

type AccountRoutes struct {
	accountController *controllers.AccountController
}

func (r *AccountRoutes) InitRoutes(groupRouter *gin.RouterGroup, accountController *controllers.AccountController) {
	r.accountController = accountController

	router := groupRouter.Group("/account")
	{
		router.Use(middleware.CheckMiddleware)
		router.GET("/balance", r.accountController.GetBalance())
		router.POST("/transfer", r.accountController.Transfer())
	}
}
