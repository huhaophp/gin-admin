package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
