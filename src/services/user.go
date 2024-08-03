package services

import (
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/forms"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func AllUsers() (users []models.User, err error) {
	_, err = db.GetDB().Select(&users, "SELECT * FROM users")
	log.Printf("%+v", users)
	return users, err
}

func Create(uf forms.UserForm) (err error) {
	newUser := models.User{
		ID:        0, // Auto incrementing as specified in ../db/db.go
		Email:     uf.Email,
		Name:      uf.Name,
		Password:  uf.Password,
		CreatedAt: 0, // Will be changed to normal value, specified in ../models/user.go
		UpdatedAt: 0, // Same
	}
	err = db.GetDB().Insert(&newUser)
	return err
}

func Login(lf forms.LoginForm) (*models.User, error) {
	var user models.User
	dbError := db.GetDB().SelectOne(&user, "SELECT * FROM users WHERE email = $1", lf.Email)

	if dbError != nil {
		return nil, dbError
	}
	log.Printf("%s and %s", user.Password, lf.Password)
	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password))
	// !!! Turns out bcrypt is slow on purpose, so don't worry:)
	if passwordError != nil {
		return nil, passwordError
	}

	return &user, nil
}
