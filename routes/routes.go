package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/http/handlers"
	"github.com/huhaophp/gin-admin/app/http/middware"
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
