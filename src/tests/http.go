package tests

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func PrintBody(c *gin.Context) {
	res, err := io.ReadAll(c.Request.Body)
	log.Printf("Request body: %s", string(res))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Message": err})
	}
	c.JSON(200, gin.H{"message": res})
	return
}
