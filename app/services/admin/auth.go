package admin

import (
	"ginapi/app/http/request/admin"
	"ginapi/app/model"
	"ginapi/tools"
	"github.com/gin-gonic/gin"
)

func NewAuthService() *authService {
	return &authService{}
}

type authService struct{}

// Login Handling user login
func (*authService) Login(c *gin.Context, r *admin.AuthLoginRequest) (error, gin.H) {
	err, user := model.GetUserByName(r.Username)
	if err != nil {
		return err, nil
	}
	token, err := tools.MakeJwtToken(user.Id)
	if err != nil {
		return err, nil
	}
	return nil, gin.H{"token": token}
}
