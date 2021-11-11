package model

import (
	"errors"
	"github.com/huhaophp/gin-admin/app/http/request"
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

// GetUsersList GetUsers
func GetUsersList(params *request.GetUsersListParam) (error, []*User, int64) {
	var (
		users []*User
		total int64
	)

	query := db.DB.Table("users")
	if params.Username != "" {
		query = query.Where("username like ?", "%"+params.Username+"%")
	}
	if params.Email != "" {
		query = query.Where("email like ?", "%"+params.Email+"%")
	}
	if params.Phone != "" {
		query = query.Where("phone like ?", "%"+params.Phone+"%")
	}

	findError := query.Limit(params.Size).Offset((params.Page - 1) * params.Size).Find(&users).Error
	countError := query.Count(&total).Error

	if findError != nil || countError != nil {
		return errors.New("user: find or count error"), nil, 0
	}

	return nil, users, total
}
