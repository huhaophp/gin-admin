package handlers

import (
	"ginapi/app/http/request/admin"
	"ginapi/app/services"
	"ginapi/tools/response"

	"github.com/gin-gonic/gin"
)

var Auth = auth{}

type auth struct{}

func (*auth) Login(ctx *gin.Context) {
	var request admin.AuthLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.Msg(ctx, err.Error(), 422)
		return
	}
	err, user := services.AdminAuth.Login(ctx, &request)
	if err != nil {
		response.Msg(ctx, err.Error(), 522)
	} else {
		response.Data(ctx, user)
	}
}
