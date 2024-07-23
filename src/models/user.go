package models

type User struct {
	ID        int64  `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Name      string `db:"name"`
	UpdatedAt int64  `db:"updated_at"`
	CreatedAt int64  `db:"created_at"`
}
