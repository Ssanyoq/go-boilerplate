package routers

import "github.com/gin-gonic/gin"

func NotFound(c *gin.Context) {
	c.JSON(404, "Error: This page does not exist (but look at that handsome 404 page tho)")
}
