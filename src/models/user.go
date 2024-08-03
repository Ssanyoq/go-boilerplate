package models

import (
	"github.com/go-gorp/gorp"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User structure is mostly used in database actions and in
// some responses. db tag specifies how corresponding fields
// should be named in database, json specifies how to output
// fields in JSON(responses and requests)
type User struct {
	ID        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
}

// PreInsert is a gorp hook that is called every time *gorp.DbMap Insert
// method is used with an instance of User.
//
// This hook adds info
// to necessary timestamps at User.CreatedAt and User.UpdatedAt
// and hashes User.Password.
//
// It is required that given password
// is not longer than 72 bytes since
// bcrypt.GenerateFromPassword method does not allow strings
// longer than that
func (u *User) PreInsert(s gorp.SqlExecutor) error {
	// Remember that this method name is necessary to maintain,
	// as it's the only one that gorp calls
	u.CreatedAt = time.Now().UnixNano()
	u.UpdatedAt = u.CreatedAt
	pswd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(pswd)
	return nil
}

// PreUpdate is a gorp hook that is called every time *gorp.DbMap Update
// method is used with an instance of User.
//
// This hook changes update time User.UpdatedAt to current server time and hashes User.Password.
//
// It is required that given password
// is not longer than 72 bytes since
// bcrypt.GenerateFromPassword method does not allow strings
// longer than that
func (u *User) PreUpdate(s gorp.SqlExecutor) error {
	u.UpdatedAt = time.Now().UnixNano()
	pswd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(pswd)
	return nil
}
