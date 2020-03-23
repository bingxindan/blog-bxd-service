package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetDetail(c *gin.Context) {
	id := c.Param("id")

	log.Printf("sssss %d", 11)

	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
