package admin

import (
	"ginapi/app/http/request/admin"
	"ginapi/app/services"
	"ginapi/tools"

	"github.com/gin-gonic/gin"
)

// Login 管理员登录 检查账号状态
func Login(ctx *gin.Context) {
	var request admin.AuthLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		tools.JsonReturnWithError(ctx, err.Error(), 422)
		return
	}
	err, user := services.AdminAuth.HandleLogin(ctx, &request)
	if err != nil {
		tools.JsonReturnWithError(ctx, err.Error(), 422)
	} else {
		tools.JsonReturnWithData(ctx, user)
	}
}

// Info 获取管理员登录信息
func Info(ctx *gin.Context) {
	err, user := services.AdminAuth.GetUssrInfo(ctx)
	if err != nil {
		tools.JsonReturnWithError(ctx, err.Error(), 422)
	} else {
		tools.JsonReturnWithData(ctx, user)
	}
}
