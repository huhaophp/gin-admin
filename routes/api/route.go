package api

import (
	"ginapi/app/http/handlers/admin"

	"github.com/gin-gonic/gin"
)

func Routers(engine *gin.Engine) {
	backend := engine.Group("api")
	{
		backend.POST("token", admin.Login)
	}
}
