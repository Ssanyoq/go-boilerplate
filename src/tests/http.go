package tests

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func PrintBody(c *gin.Context) {
	res, _ := io.ReadAll(c.Request.Body)
	log.Printf("Request body: %s", string(res))
	return
}
