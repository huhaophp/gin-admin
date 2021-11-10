package admin

import (
	"ginapi/app/http/request/admin"
	"ginapi/app/services"
	"ginapi/tools/response"

	"github.com/gin-gonic/gin"
)

// Login 管理员登录 检查账号状态
func Login(ctx *gin.Context) {
	var request admin.AuthLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.Message(ctx, err.Error(), services.ParamsErrCode)
		return
	}
	err, user := services.AdminAuth.Login(ctx, &request)
	if err != nil {
		response.Message(ctx, err.Error(), services.ParamsErrCode)
	} else {
		response.Data(ctx, user)
	}
}
