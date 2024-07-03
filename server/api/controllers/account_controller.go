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
	type TransferRequest struct {
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}
	return func(c *gin.Context) {
		userName := c.GetString("username")

		transferRequest := TransferRequest{}

		if err := c.ShouldBindJSON(&transferRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}

		if transferRequest.Amount <= 0 || transferRequest.To == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
		}

		if userName == transferRequest.To {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "You cannot send money to yourself",
			})
			return
		}

		accDetails, err := a.accountService.GetBalanceByUserName(userName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		if accDetails.Balance < transferRequest.Amount {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Insufficient balance",
			})
			return
		}

		err = a.accountService.TransferAmount(userName, transferRequest.To, transferRequest.Amount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message ": "Amount transferred successfully",
		})
	}
}
