package middleware

import (
	"net/http"
	"server/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	token := strings.Split(header, " ")[1]

	if token == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	claims, err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}
