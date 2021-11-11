package routes

import (
	"ginapi/app/http/handlers"
	"ginapi/app/http/middware"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) *gin.Engine {
	engine.POST("login", handlers.Auth.Login)
	engine.Use(middware.JWTAuthMiddleware())
	{
		engine.GET("users/list", handlers.User.List)
		engine.POST("users/store", handlers.User.Store)
		engine.POST("users/delete", handlers.User.Delete)
	}
	return engine
}
