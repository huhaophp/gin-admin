package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/http/request"
	"github.com/huhaophp/gin-admin/app/model"
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
	if err, list, total := model.GetUsersList(&req); err != nil {
		response.Msg(ctx, err.Error(), 422)
	} else {
		response.Data(ctx, response.PageResult{
			List:  list,
			Total: total,
			Page:  req.Page,
			Size:  req.Size,
		})
	}
}

func (*user) Store(ctx *gin.Context) {
}

func (*user) Delete(ctx *gin.Context) {
}
