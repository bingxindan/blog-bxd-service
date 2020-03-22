package bootstrap

import (
	"blog-bxd-service/gateways/order/routes"
	"github.com/gin-gonic/gin"
)

func StartOrderServer(init *gin.Engine) {
	routers.InitRouter(init)
}
