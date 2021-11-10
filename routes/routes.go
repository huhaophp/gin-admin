package routes

import (
	"ginapi/routes/admin"
	"ginapi/routes/api"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) *gin.Engine {
	api.Routers(engine)
	admin.Routers(engine)
	return engine
}
