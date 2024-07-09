package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/controllers"
)

func (r routes) addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", controllers.AllUsers)
}
