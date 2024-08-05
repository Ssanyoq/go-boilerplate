package main

import (
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/routers"
	"os"
)

func StartServer() {
	db.Init()
	hostname := os.Getenv("HOST_NAME")
	router := routers.NewRoutes()
	router.Run(hostname)
}

func main() {
	StartServer()
}
