package routers

import "github.com/gin-gonic/gin"

type routes struct {
	router *gin.Engine
}

// NewRoutes Creates routes instance and adds necessary router
// groups to it. Attempt to make scalable routing
func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group("/v1")
	r.addUserRoutes(v1)
	r.AddShopRoutes(v1)
	return r
}

func (r routes) Run(addr string) error {
	return r.router.Run(addr)
}
