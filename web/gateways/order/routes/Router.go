package routers

import (
	"blog-bxd-service/gateways/order/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/bxd-order-middleware")

	api.GET("/order/list/:orderId", controller.GetList)
}
