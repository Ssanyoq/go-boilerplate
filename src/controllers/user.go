package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/forms"
	"github.com/ssanyoq/go-boilerplate/src/services"
	"log"
	"net/http"
)

func AllUsers(c *gin.Context) {
	result, err := services.AllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Database is ill today:("})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": result})
}

func OneUser(c *gin.Context) {

}

func Create(c *gin.Context) {
	var userForm forms.UserForm
	bindingError := c.ShouldBindJSON(&userForm)
	if bindingError != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": bindingError.Error()})
		return
	}
	validationError := forms.ValidateUserForm(userForm)
	if validationError != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": validationError})
		log.Print("Invalid thingy")
		return
	}
	res, err := services.Create(userForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err})
		log.Print("Bad dadabada")
		log.Printf("%s", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": res})
}
