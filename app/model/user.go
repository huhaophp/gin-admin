package model

import (
	"database/sql"
	"ginapi/support/db"
	"time"
)

type (
	User struct {
		Id        int64     `db:"id" json:"id" form:"id"`
		Username  string    `db:"username" json:"username" form:"username"`
		Password  string    `db:"password" json:"-" form:"password"`
		CreatedAt time.Time `db:"created_at" json:"created_at" form:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
	}
)

// GetUserByName by username get user.
func GetUserByName(username string) (error, *User) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = ? limit 1", username)
	if err != sql.ErrNoRows {
		return err, nil
	} else {
		return nil, &user
	}
}
