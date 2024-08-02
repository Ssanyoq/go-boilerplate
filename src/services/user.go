package services

import (
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/forms"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"log"
)

func AllUsers() (users []models.User, err error) {
	_, err = db.GetDB().Select(&users, "SELECT * FROM users")
	log.Printf("%+v", users)
	return users, err
}

func Create(uf forms.UserForm) (err error) {
	newUser := models.User{
		ID:        0,
		Email:     uf.Email,
		Name:      uf.Name,
		Password:  uf.Password,
		CreatedAt: 0,
		UpdatedAt: 0,
	}
	err = db.GetDB().Insert(&newUser)
	return err
}
