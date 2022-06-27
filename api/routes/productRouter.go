package routes

import (
	controller "GOLANG/api/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/product/all", controller.AllProduct())
	incomingRoutes.POST("/product/id", controller.ProductById())
	incomingRoutes.POST("/product/popular", controller.PopularProduct())
	incomingRoutes.POST("/product/price", controller.ProductByPrice())
}
