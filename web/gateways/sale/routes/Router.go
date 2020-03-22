package routers

import (
	"blog-bxd-service/gateways/sale/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/bxd-sale-middleware")

	api.GET("/goods/detail/:id", controller.GetDetail)
}
