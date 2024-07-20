package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"log"
	"net/http"
)

func AllUsers(c *gin.Context) {
	result, err := models.AllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Database is ill today:("})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": result})
}
