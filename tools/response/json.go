package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

func Data(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"code": 0,
		"data": data,
	})
}

func Msg(ctx *gin.Context, msg string, code int) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
		"data": nil,
	})
}
