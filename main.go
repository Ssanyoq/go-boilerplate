package main

import "github.com/ssanyoq/go-boilerplate/src/routers"

func main() {
	router := routers.NewRoutes()
	router.Run("localhost:8000")

}
