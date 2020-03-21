package bootstrap

import (
	"bxd-middleware-service/gateways/order/routes"
	"github.com/gin-gonic/gin"
)

func StartOrderServer(init *gin.Engine) {
	routers.InitRouter(init)
}
