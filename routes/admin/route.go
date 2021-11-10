package admin

import (
	"ginapi/app/http/handlers/admin"
	"ginapi/app/http/middware"

	"github.com/gin-gonic/gin"
)

func Routers(engine *gin.Engine) {
	backend := engine.Group("backend")
	{
		backend.POST("login", admin.Login)
		backend.Use(middware.JWTAuthMiddleware())
		{
			backend.GET("userinfo", admin.Info)
		}
	}
}
