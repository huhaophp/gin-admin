package handlers

import (
	"ginapi/app/model"
	"ginapi/tools/response"
	"github.com/gin-gonic/gin"
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
