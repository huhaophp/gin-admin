package handlers

import (
	"github.com/huhaophp/gin-admin/app/http/request"
	"github.com/huhaophp/gin-admin/app/services"
	"github.com/huhaophp/gin-admin/tools/response"

	"github.com/gin-gonic/gin"
)

var Auth = &auth{}

type auth struct{}

func (*auth) Login(ctx *gin.Context) {
	var req request.AuthLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Msg(ctx, err.Error(), 422)
		return
	}
	err, user := services.AuthService.Login(ctx, &req)
	if err != nil {
		response.Msg(ctx, err.Error(), 422)
	} else {
		response.Data(ctx, user)
	}
}
