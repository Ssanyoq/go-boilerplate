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

func Create(uf forms.UserForm) (res int64, err error) {
	//newUser := models.User{
	//	ID:        0,
	//	Email:     uf.Email,
	//	Name:      uf.Name,
	//	Password:  uf.Password,
	//	CreatedAt: 0,
	//	UpdatedAt: 0,
	//}
	//err = db.GetDB().Insert(&newUser)
	result, err := db.GetDB().Exec("insert into users (id, name, email, password, updated_at, created_at) values"+
		" ($1, $2, $3, $4, $5, $6) returning id",
		0, uf.Email, uf.Name, uf.Password, 0, 0)
	res, _ = result.LastInsertId()
	log.Printf("Result is %s", res)
	return res, err
}
