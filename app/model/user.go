package model

import (
	"database/sql"
	"github.com/huhaophp/gin-admin/support/db"
	"time"
)

type (
	User struct {
		Id        int64     `db:"id" json:"id" form:"id"`
		Username  string    `db:"username" json:"username" form:"username"`
		Phone     string    `db:"phone" json:"phone" form:"phone"`
		Email     string    `db:"email" json:"email" form:"email"`
		Password  string    `db:"password" json:"-" form:"password"`
		CreatedAt time.Time `db:"created_at" json:"created_at" form:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
	}
)

// GetUserByUsername by username get user.
func GetUserByUsername(username string) (error, *User) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = ? limit 1", username)
	if err != nil && err != sql.ErrNoRows {
		return err, nil
	}
	return nil, &user
}

// GetUsers GetUsers
func GetUsers() (error, []*User) {
	var users []*User
	err := db.DB.Select(&users, "SELECT * FROM users limit 20 offset 0")
	if err != nil && err != sql.ErrNoRows {
		return err, nil
	}
	return nil, users
}
