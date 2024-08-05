package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/middlewares"
	"log"
)

type routes struct {
	router *gin.Engine
}

// NewRoutes Creates routes instance and adds necessary router
// groups to it. Attempt to make scalable routing
func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(middlewares.CORSMiddleware())
	v1 := r.router.Group("/v1")
	r.addUserRoutes(v1)
	r.AddShopRoutes(v1)
	r.router.NoRoute(NotFound)
	return r
}

func (r routes) Run(addr string) error {
	log.Printf("Go to http://%s to use your stuff!", addr)
	return r.router.Run(addr)
}

func (r routes) GetEngine() *gin.Engine {
	return r.router
}
