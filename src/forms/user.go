package forms

import (
	"errors"
	"github.com/ssanyoq/go-boilerplate/src/validators"
	"strings"
)

type UserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateUserForm(f UserForm) error {
	errorMsg := make([]string, 0)

	if !validators.Name(f.Name) {
		errorMsg = append(errorMsg, "Name does not match requirements")
	}
	if !validators.Password(f.Password) {
		errorMsg = append(errorMsg, "Password does not match requirements")
	}
	if !validators.Email(f.Email) {
		errorMsg = append(errorMsg, "Email is not an email")
	}
	if len(errorMsg) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorMsg, "; "))
}
