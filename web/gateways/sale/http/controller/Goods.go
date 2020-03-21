package controller

import (
	"bxd-middleware-service/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDetail(c *gin.Context) {
	id := c.Param("id")

	log.Infof("233333s")

	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
