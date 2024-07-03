package controllers

import (
	"net/http"
	"server/api/services"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *services.AccountServices
}

func (a *AccountController) InitController(accountService *services.AccountServices) *AccountController {
	a.accountService = accountService
	return a
}

func (a *AccountController) GetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.GetString("username")
		accDetails, err := a.accountService.GetBalanceByUserName(userName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		c.JSON(200, gin.H{
			"balance": accDetails.Balance,
		})
	}
}

func (a *AccountController) Transfer() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
