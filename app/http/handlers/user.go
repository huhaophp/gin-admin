package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/model"
	"github.com/huhaophp/gin-admin/tools/response"
)

var User = user{}

type user struct{}

func (*user) List(ctx *gin.Context) {
	_, users := model.GetUsers()
	response.Data(ctx, users)
}

func (*user) Store(ctx *gin.Context) {
}

func (*user) Delete(ctx *gin.Context) {
}
