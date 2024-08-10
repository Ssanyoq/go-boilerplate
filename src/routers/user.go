package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/controllers"
	"github.com/ssanyoq/go-boilerplate/src/middlewares"
	"github.com/ssanyoq/go-boilerplate/src/tests"
)

func (r routes) addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", middlewares.RequireAuth, controllers.AllUsers)
	users.POST("/", controllers.Create)
	users.PUT("/", tests.PrintBody) // FOR TESTS
	users.POST("/login", controllers.Login)

}
