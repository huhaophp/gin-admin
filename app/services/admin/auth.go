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

// HandleLogin Handling user login
func (a *authService) HandleLogin(c *gin.Context, r *admin.AuthLoginRequest) (error, gin.H) {
	err, user := model.GetUserByName(r.Name)
	if err != nil {
		return err, nil
	}
	token, err := tools.MakeJwtToken(user.Id)
	if err != nil {
		return err, nil
	}
	return nil, gin.H{"token": token}
}

// GetUssrInfo Handling user login
func (a *authService) GetUssrInfo(c *gin.Context) (error, *model.User) {
	//uid, _ := c.Get("jwt_uid")
	//err, user := dao.User.GetUserById(uid.(int64))
	//if err != nil {
	//	return err, nil
	//}
	//return nil, user
	return nil, nil
}
