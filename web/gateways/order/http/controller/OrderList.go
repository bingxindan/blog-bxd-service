package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context) {
	orderId := c.Param("orderId")

	c.JSON(http.StatusOK, gin.H{
		"data": orderId,
	})
}
