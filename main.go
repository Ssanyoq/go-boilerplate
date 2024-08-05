package main

import (
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/routers"
	"os"
)

func main() {

	db.Init()

	hostname := os.Getenv("HOST_NAME")
	router := routers.NewRoutes()
	router.Run(hostname)

}
