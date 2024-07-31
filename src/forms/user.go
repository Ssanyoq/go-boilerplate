package forms

import (
	"errors"
	"github.com/ssanyoq/go-boilerplate/src/validators"
	"log"
	"strings"
)

type UserForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateUserForm(f UserForm) error {
	errorMsg := make([]string, 0)

	if !validators.Name(f.Name) {
		errorMsg = append(errorMsg, "Name does not match requirements")
		log.Printf("%s is not a name", f.Name)
	}
	if !validators.Password(f.Password) {
		errorMsg = append(errorMsg, "Password does not match requirements")
		log.Printf("%s is not a passwd", f.Password)
	}
	if !validators.Email(f.Email) {
		errorMsg = append(errorMsg, "Email is not an email")
		log.Printf("%s is not an email", f.Email)
	}
	if len(errorMsg) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorMsg, "; "))
}
