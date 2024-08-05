package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/routers"
	"github.com/ssanyoq/go-boilerplate/src/utility"
	"os"
)

func StartServer() *gin.Engine {
	utility.SetupEnv()
	db.Init()
	hostname := os.Getenv("HOST_NAME")
	router := routers.NewRoutes()
	router.Run(hostname)
	return router.GetEngine()
}
