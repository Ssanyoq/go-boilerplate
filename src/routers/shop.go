package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/controllers"
)

func (r routes) AddShopRoutes(rg *gin.RouterGroup) {
	shop := rg.Group("/shop")
	shop.GET("/", controllers.AllItems)
}
