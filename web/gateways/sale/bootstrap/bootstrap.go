package bootstrap

import (
	"blog-bxd-service/gateways/sale/routes"
	"github.com/gin-gonic/gin"
)

func StartSaleServer(init *gin.Engine) {
	routers.InitRouter(init)
}
