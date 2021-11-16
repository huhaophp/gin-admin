package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/http/request"
	"github.com/huhaophp/gin-admin/app/services"
	"github.com/huhaophp/gin-admin/tools/response"
)

var User = user{}

type user struct{}

func (*user) List(ctx *gin.Context) {
	var req request.GetUsersListParam
	if err := ctx.BindQuery(&req); err != nil {
		response.Msg(ctx, err.Error(), 422)
		return
	}
	err, result := services.UserService.GetUsersList(ctx, &req)
	if err != nil {
		response.Msg(ctx, err.Error(), 500)
		return
	}
	response.Data(ctx, result)
}

func (*user) Store(ctx *gin.Context) {
	response.Data(ctx, nil)
}

func (*user) Delete(ctx *gin.Context) {
}
