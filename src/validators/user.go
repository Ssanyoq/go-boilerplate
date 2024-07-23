package validators

import (
	"net/mail"
	"regexp"
)

// Email validator checks if given string matches a correct
// email address
//
// See: mail.ParseAddress()
func Email(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

// Password validator checks if password is between 8 and 255
// characters long
func Password(p string) bool {
	return len(p) >= 8 && len(p) < 256 // Temporary (probably)
}

// Name validator checks if given string is from 2 to 20 letters
// long and is a string
// (without any other symbols like spaces, numbers, etc.)
func Name(n string) bool {
	isString, _ := regexp.Match(`[:alpha]+`, []byte(n))
	return len(n) >= 2 && len(n) <= 20 && isString
}
