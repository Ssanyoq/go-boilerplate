package services

import (
	"github.com/ssanyoq/go-boilerplate/src/db"
	"github.com/ssanyoq/go-boilerplate/src/forms"
	"github.com/ssanyoq/go-boilerplate/src/models"
)

func AllUsers() (users []models.User, err error) {
	_, err = db.GetDB().Select(&users, "SELECT * FROM users")
	return users, err
}

func Create(uf forms.UserForm) (id int64, err error) {
	row := db.GetDB().QueryRow("INSERT INTO users(name, password, email) VALUES($1, $2, $3) RETURNING id",
		uf.Name, uf.Password, uf.Email)
	err = row.Scan(&id)
	return id, err
}
