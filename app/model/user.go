package model

import (
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
	err := db.DB.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	return nil, &user
}

// GetUsersList GetUsersList
func GetUsersList(params map[string]interface{}) (error, []*User, int64) {
	var total int64
	var users []*User

	query := db.DB.Table("users")
	if username, ok := params["username"]; ok && username != "" {
		query = query.Where("username like ?", username.(string)+"%")
	}
	if email, ok := params["email"]; ok && email != "" {
		query = query.Where("email like ?", email.(string)+"%")
	}
	if phone, ok := params["phone"]; ok && phone != "" {
		query = query.Where("phone like ?", phone.(string)+"%")
	}

	query.Limit(params["size"].(int)).Offset((params["page"].(int) - 1) * params["size"].(int)).Find(&users)
	query.Count(&total)

	return nil, users, total
}
