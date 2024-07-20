package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/routers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	setModeFromEnv()
	db.Init()

	hostname := os.Getenv("HOST_NAME")
	router := routers.NewRoutes()

	log.Printf("Go to http://%s to use your stuff!", hostname)
	router.Run(hostname)

}

func setModeFromEnv() {
	switch os.Getenv("ENV") {
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	}
}
