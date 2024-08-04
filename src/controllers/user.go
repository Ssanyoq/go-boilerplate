package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssanyoq/go-boilerplate/src/forms"
	"github.com/ssanyoq/go-boilerplate/src/services"
	"github.com/ssanyoq/go-boilerplate/src/utility"
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

func Login(c *gin.Context) {
	var loginForm forms.LoginForm
	bindingError := c.ShouldBindJSON(&loginForm)
	if bindingError != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": bindingError.Error()})
		return
	}
	user, err := services.Login(loginForm)
	if err != nil {
		utility.ClearTokens(c) // grrr bad user grrr
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	err = utility.SetTokens(c, *user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": user.ID, "name": user.Name})

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
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": validationError.Error()})
		log.Print("Invalid thingy")
		return
	}
	err := services.Create(userForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		log.Printf("%s", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
