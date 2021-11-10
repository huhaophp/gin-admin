package model

import (
	"ginapi/support/db"
	"time"
)

type (
	User struct {
		Id        int64     `db:"id" json:"id" form:"id"`
		Name      string    `db:"name" json:"name" form:"name"`
		Email     string    `db:"email" json:"email" form:"email"`
		Password  string    `db:"password" json:"-" form:"password"`
		CreatedAt time.Time `db:"created_at" json:"created_at" form:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
	}
)

// GetUserByName by name get user.
func GetUserByName(name string) (error, *User) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE name = ? limit 1", name)
	if err != nil {
		return err, nil
	}
	return nil, &user
}

// GetUserByEmail by email get user.
func GetUserByEmail(email string) (error, *User) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE email = ? limit 1", email)
	if err != nil {
		return err, nil
	}
	return nil, &user
}
