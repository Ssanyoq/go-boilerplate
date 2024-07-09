package controllers

import "github.com/gin-gonic/gin"

func AllUsers(c *gin.Context) {
	c.IndentedJSON(200, "Hello world!")
}
