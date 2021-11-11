package services

import (
	"errors"
	"github.com/huhaophp/gin-admin/app/http/request"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/model"
	"github.com/huhaophp/gin-admin/tools"
)

var Auth = &authService{}

type authService struct{}

// Login Handling user login
func (*authService) Login(ctx *gin.Context, request *request.AuthLoginRequest) (error, map[string]string) {
	err, user := model.GetUserByUsername(request.Username)
	if err != nil || user.Password != tools.Md5(request.Password) {
		return errors.New("username or password error"), nil
	}
	token, err := tools.MakeJwtToken(user.Id)
	if err != nil {
		return errors.New("make token error"), nil
	} else {
		return nil, map[string]string{"token": token}
	}
}
